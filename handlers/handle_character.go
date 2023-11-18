package handlers

import (
	"encoding/json"
	"github.com/banaaron/tekken-backend/database"
	"net/http"
)

func HandleCharacter(writer http.ResponseWriter, _ *http.Request) {
	var characters []database.Character
	var err error

	EnableCors(&writer)

	characters, err = database.GetCharacters()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(characters)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
