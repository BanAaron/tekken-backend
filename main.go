package main

import (
	"database/sql"
	"fmt"
	"github.com/banaaron/tekken-backend/database"
	"github.com/banaaron/tekken-backend/handlers"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func init() {
	err := database.InitDb()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var err error
	githubURL := "https://github.com/aarontbarratt/tekken-backend#tekken-backend"
	port := os.Getenv("PORT")
	fmt.Printf("http://localhost:%s\n", port)

	server := http.NewServeMux()
	server.HandleFunc("/", handlers.HandleRoot)
	server.Handle("/api/help", http.RedirectHandler(githubURL, http.StatusSeeOther))
	server.HandleFunc("/api/teapot", handlers.HandleTeapot)
	server.HandleFunc("/api/character", handlers.HandleCharacter)
	err = http.ListenAndServe(":"+port, server)
	if err != nil {
		log.Fatal(err)
	}

	defer func(Db *sql.DB) {
		err := Db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(database.Db)
}
