package database

// https://go.dev/doc/tutorial/database-access

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const driver = "postgres"

var DbConnectionString ConnectionString

// CheckDatabaseConnection connects to and pings the server to make sure the connection is working
func CheckDatabaseConnection() error {
	db, err := sql.Open(driver, DbConnectionString.Get())
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// GetCharacters requests all the characters from the database
func GetCharacters() ([]Character, error) {
	db, err := sql.Open(driver, DbConnectionString.Get())
	if err != nil {
		return nil, err
	}
	var characters []Character
	rows, err := db.Query(`
		select id, short_name, long_name, fighting_style, nationality, height, weight, gender
		from characters
		order by id
	`)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			println(fmt.Errorf("failed to close databse connection: %s", err))
		}
	}()

	for rows.Next() {
		var (
			id            int
			shortName     string
			longName      string
			fightingStyle string
			nationality   string
			height        int
			weight        int
			gender        string
		)

		err := rows.Scan(&id, &shortName, &longName, &fightingStyle, &nationality, &height, &weight, &gender)
		if err != nil {
			return nil, err
		}
		newCharacter := NewCharacter(id, shortName, longName, fightingStyle, nationality, height, weight, gender)
		characters = append(characters, newCharacter)
	}
	if err != nil {
		return nil, err
	}
	return characters, nil
}
