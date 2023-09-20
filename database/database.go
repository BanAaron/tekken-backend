package database

// https://go.dev/doc/tutorial/database-access

import (
	"database/sql"
	"fmt"

	"aaronbarratt.dev/go/tekken-backend/utils"
	_ "github.com/lib/pq"
)

const driver = "postgres"

var DbConnectionString ConnectionString

// CheckDatabaseConnection connects to and pings the server to make sure the connection is working
func CheckDatabaseConnection() (err error) {
	db, err := sql.Open(driver, DbConnectionString.Get())
	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("failed to close database connection"), err)
		}
	}()
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func GetCharactersWithId() (characterWithIds []CharacterWithId, err error) {
	var rows *sql.Rows

	db, err := sql.Open(driver, DbConnectionString.Get())
	if err != nil {
		return
	}
	defer func() {
		err := db.Close()
		if err != nil {
			println(fmt.Errorf("failed to close databse connection: %s", err))
		}
	}()

	rows, err = db.Query("select id, short_name from characters order by short_name")
	if err != nil {
		return
	}

	for rows.Next() {
		var (
			id        int
			shortName string
		)

		err := rows.Scan(&id, &shortName)
		if err != nil {
			fmt.Println(err)
		}
		characterWithIds = append(characterWithIds, CharacterWithId{
			Id:        id,
			ShortName: shortName,
		})
	}

	return
}

// GetCharacters requests all the characters from the database
func GetCharacters(name string) (characters []Character, err error) {
	var query string
	var rows *sql.Rows
	util.ToTitleCase(&name)

	db, err := sql.Open(driver, DbConnectionString.Get())
	if err != nil {
		return nil, err
	}

	if name == "" {
		query = `
			select id, short_name, long_name, fighting_style, nationality, height, weight, gender
			from characters
			order by id
		`
		rows, err = db.Query(query)
	} else {
		query = `
			select id, short_name, long_name, fighting_style, nationality, height, weight, gender
			from characters
			where short_name = $1
		`
		rows, err = db.Query(query, name)
	}
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
