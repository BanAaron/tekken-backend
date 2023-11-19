package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var Db *sql.DB

const driver = "postgres"

func getEnvironmentVariables() (username string, password string, host string, dbName string, port int, err error) {
	hostname, err := os.Hostname()
	if err != nil {
		return
	}
	if hostname == "aaron" {
		err = godotenv.Load(".env")
		if err != nil {
			return
		}
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

func InitDb() (err error) {
	username, password, host, dbName, port, err := getEnvironmentVariables()
	if err != nil {
		return
	}

	connectionString := NewConnectionString(username, password, host, port, dbName).Get()
	Db, err = sql.Open(driver, connectionString)
	if err != nil {
		return
	}
	fmt.Println("db connection successful")

	err = Db.Ping()
	if err != nil {
		message := fmt.Errorf("db.ping: %s", err)
		return message
	}
	fmt.Println("database ping successful")
	return nil
}
