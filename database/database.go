package database

// https://go.dev/doc/tutorial/database-access

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

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

func GetCharacters(writer http.ResponseWriter, _ *http.Request) {
	db, err := sql.Open(driver, DbConnectionString.Get())
	if err != nil {
		http.Error(writer, "Unable to connect to database", http.StatusInternalServerError)
	}
	var characters []Character
	rows, err := db.Query(`
		select id, short_name, long_name, fighting_style, nationality, height, weight, gender
		from characters
		order by id
	`)
	if err != nil {
		http.Error(writer, "Failed to query database", http.StatusInternalServerError)
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
			http.Error(writer, "Failed to read data from rows", http.StatusInternalServerError)
		}
		newCharacter := NewCharacter(id, shortName, longName, fightingStyle, nationality, height, weight, gender)
		characters = append(characters, newCharacter)
	}
	err = json.NewEncoder(writer).Encode(characters)
	if err != nil {
		http.Error(writer, "Failed to read data from rows", http.StatusInternalServerError)
	}
}

func GetCharacter(characterShortName string, connectionString ConnectionString) (character *Character, err error) {
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
	db, err := sql.Open(driver, connectionString.Get())
	if err != nil {
		return nil, err
	}
	row := db.QueryRow(`
		select id, short_name, long_name, fighting_style, nationality, height, weight, gender
		from characters 
		where short_name = $1`, characterShortName)
	err = row.Scan(
		&id,
		&shortName,
		&longName,
		&fightingStyle,
		&nationality,
		&height,
		&weight,
		&gender,
	)
	if err != nil {
		return nil, err
	}
	newCharacter := NewCharacter(id, shortName, longName, fightingStyle, nationality, height, weight, gender)

	return &newCharacter, nil
}
