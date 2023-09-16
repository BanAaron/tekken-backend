package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"aaronbarratt.dev/go/tekken-backend/database"
	util "aaronbarratt.dev/go/tekken-backend/utils"
	_ "github.com/lib/pq"
)

func main() {
	var err error
	err = util.LoadDotEnv()
	if err != nil {
		log.Fatal(err)
	}
	database.DbConnectionString = database.NewConnectionString(util.GetEnvVariables())

	// check that the connection string is working with the database
	err = database.CheckDatabaseConnection()
	util.CheckError(err)
	fmt.Println("Database connection success!")

	// create the server
	server := http.NewServeMux()
	// create routes
	server.Handle("/", http.RedirectHandler("https://github.com/aarontbarratt/tekken-backend", http.StatusSeeOther))
	server.HandleFunc("/teapot", util.HandleTeapot)
	server.HandleFunc("/api/characters", handleCharacters)
	// start the server
	err = http.ListenAndServe(":8888", server)
	util.CheckError(err)
}

func handleCharacters(writer http.ResponseWriter, _ *http.Request) {
	characters, err := database.GetCharacters()
	if err != nil {
		writer.Header().Set("Error", err.Error())
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	// encodes the characters array of structs into json and writes it to writer
	err = json.NewEncoder(writer).Encode(characters)
	if err != nil {
		writer.Header().Set("Error", err.Error())
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
