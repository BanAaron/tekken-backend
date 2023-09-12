package database

import (
	"database/sql"
	"fmt"

	util "aaronbarratt.dev/go/tekken-backend/utils"
	_ "github.com/lib/pq"
)

type Character struct {
	Id        int
	ShortName string
}

type ConnectionString struct {
	username string
	password string
	host     string
	port     int
	dbname   string
}

func (cs ConnectionString) String() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s",
		cs.username,
		"[PASSWORD]",
		cs.host,
		cs.port,
		"[DBNAME]",
	)
}

func (cs ConnectionString) Get() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s",
		cs.username,
		cs.password,
		cs.host,
		cs.port,
		cs.dbname)
}

func NewConnectionString(
	username string,
	password string,
	host string,
	port int,
	dbname string,
) ConnectionString {
	return ConnectionString{
		username: username,
		password: password,
		host:     host,
		port:     port,
		dbname:   dbname,
	}
}

func GetCharacters(db *sql.DB) (characters []Character) {
	res, err := db.Query("select id, short_name from characters")
	util.CheckError(err)

	defer func(res *sql.Rows) {
		err := res.Close()
		util.CheckError(err)
	}(res)

	for res.Next() {
		var id int
		var shortName string

		err := res.Scan(&id, &shortName)
		util.CheckError(err)

		characters = append(characters, Character{
			Id:        id,
			ShortName: shortName,
		})

	}
	return characters
}
