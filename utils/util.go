package util

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// CheckError checks if there is an error. If there is it will throw a fatal error
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func LoadDotEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("unable to load .env: %s", err)
	}
	return nil
}

func GetEnvVariables() (username string, password string, host string, port int, dbname string) {
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

func HandleTeapot(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusTeapot)
	_, err := writer.Write([]byte("I am a teapot"))
	if err != nil {
		// Handle the error here
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
	}
}
