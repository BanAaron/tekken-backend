package main

import (
	"fmt"
	"log"
	"net/http"

	"aaronbarratt.dev/go/tekken-backend/database"
	"aaronbarratt.dev/go/tekken-backend/handlers"
	"aaronbarratt.dev/go/tekken-backend/utils"
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
	server.HandleFunc("/teapot", handlers.HandleTeapot)
	server.HandleFunc("/api/character", handlers.HandleCharacter)
	// start the server
	err = http.ListenAndServe(":8888", server)
	util.CheckError(err)
}
