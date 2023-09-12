package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"aaronbarratt.dev/go/tekken-backend/database"
	util "aaronbarratt.dev/go/tekken-backend/utils"
	_ "github.com/lib/pq"
)

func getEnvVariables() (username string, password string, host string, port int, dbname string) {
	portString := os.Getenv("DB_PORT")
	portInt, err := strconv.Atoi(portString)
	if err != nil {
		panic(fmt.Errorf("unable to parse port to int: %s", err))
	}
	port = portInt

	username = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	host = os.Getenv("DB_HOST")
	dbname = os.Getenv("DB_DATABASE_NAME")
	return username, password, host, port, dbname
}

func main() {
	util.LoadDotEnv()
	const driver = "postgres"

	connectionString := database.NewConnectionString(getEnvVariables())
	db, err := sql.Open(driver, connectionString.Get())
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	characters := database.GetCharacters(db)
	for _, character := range characters {
		fmt.Println(character)
	}
}
