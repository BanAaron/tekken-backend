package main

import (
	"fmt"
	"net/http"

	"aaronbarratt.dev/go/tekken-backend/database"
	util "aaronbarratt.dev/go/tekken-backend/utils"
	_ "github.com/lib/pq"
)

func main() {
	util.LoadDotEnv()
	database.DbConnectionString = database.NewConnectionString(util.GetEnvVariables())

	jin, err := database.GetCharacter("Jin", database.DbConnectionString)
	if err != nil {
		fmt.Println(fmt.Errorf("%s", err))
	} else {
		fmt.Println(jin)
	}

	// check that the connection string is working with the database
	err = database.CheckDatabaseConnection()
	util.CheckError(err)

	// create the server
	server := http.NewServeMux()
	// create routes
	server.Handle("/", http.RedirectHandler("https://github.com/aarontbarratt/tekken-backend", http.StatusSeeOther))
	server.HandleFunc("/teapot", util.HandleTeapot)
	server.HandleFunc("/api/characters", database.GetCharacters)
	// start the server
	err = http.ListenAndServe(":8888", server)
	util.CheckError(err)
}
