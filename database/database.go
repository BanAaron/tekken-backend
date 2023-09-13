package database

// https://go.dev/doc/tutorial/database-access

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetCharacters(db *sql.DB) (characters []Character, err error) {
	rows, err := db.Query(`
		select id, short_name, long_name, fighting_style, nationality, height, weight, gender
		from characters
		order by id
	`)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(fmt.Errorf("%s%", err))
		}
	}(rows)

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
	return characters, nil
}

func GetCharacter(characterShortName string, db *sql.DB) (character *Character, err error) {
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
