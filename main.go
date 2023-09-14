package main

import (
	"fmt"
	"os"
	"strconv"

	"aaronbarratt.dev/go/tekken-backend/database"
	util "aaronbarratt.dev/go/tekken-backend/utils"
	_ "github.com/lib/pq"
)

func main() {
	util.LoadDotEnv()
	connectionString := database.NewConnectionString(getEnvVariables())

	characters, err := database.GetCharacters(connectionString)
	if err != nil {
		fmt.Println("failed to get characters", err)
	} else {
		for _, character := range characters {
			fmt.Println(character)
		}
	}

	jin, err := database.GetCharacter("Jin", connectionString)
	if err != nil {
		panic(err)
	}
	fmt.Println(jin)
}

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
