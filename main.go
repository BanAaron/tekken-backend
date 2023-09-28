package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/banaaron/tekken-backend/database"
	"github.com/banaaron/tekken-backend/handlers"
	"github.com/banaaron/tekken-backend/utils"
)

func main() {
	var err error
	err = util.LoadDotEnv()
	if err != nil {
		log.Fatal(err)
	}

	database.DbConnectionString, err = database.NewConnectionString()
	if err != nil {
		log.Fatal(err)
	}
	handlers.FrontendHost = "http://localhost:5173"
	// check that the connection string is working with the database
	err = database.CheckDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connection success!")

	server := http.NewServeMux()
	server.Handle("/api/help", http.RedirectHandler("https://github.com/aarontbarratt/tekken-backend#tekken-backend", http.StatusSeeOther))
	server.HandleFunc("/api/character", handlers.HandleCharacter)
	server.HandleFunc("/api/characterWithId", handlers.HandleCharacterWithId)
	server.HandleFunc("/api/teapot", handlers.HandleTeapot)
	err = http.ListenAndServe(":8888", server)
	if err != nil {
		log.Fatal(err)
	}
}
