package main

import (
	"database/sql"
	"fmt"
	"github.com/banaaron/tekken-backend/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
)

func getEnvironmentVariables() (username string, password string, host string, dbName string, port int, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		return
	}
	username = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	host = os.Getenv("DB_HOST")
	dbName = os.Getenv("DB_DATABASE_NAME")
	port, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return
	}
	return
}

func main() {
	// load .env
	const driver = "postgres"
	username, password, host, dbName, port, err := getEnvironmentVariables()
	if err != nil {
		log.Fatal(err)
	}
	// connect to database
	// 1. make connection string
	connectionString := database.NewConnectionString(username, password, host, port, dbName).Get()
	// 2. connect to database
	db, err := sql.Open(driver, connectionString)
	if err != nil {
		log.Fatal(err)
	}
	// 3. defer close
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	// 4. ping db to make sure connection was successful
	err = db.Ping()
	if err != nil {
		message := fmt.Sprintf("db.ping: %s", err)
		fmt.Println(message)
	}
	// 5. query
	rows, err := db.Query("select id, short_name from characters")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var (
			id        int
			shortName string
		)
		err := rows.Scan(&id, &shortName)
		if err != nil {
			log.Fatal(err)
		}
		result := fmt.Sprintf("%d %s", id, shortName)
		fmt.Println(result)
	}
}
