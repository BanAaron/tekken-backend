package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"aaronbarratt.dev/go/tekken-backend/database"
)

func HandleCharacter(writer http.ResponseWriter, request *http.Request) {
	var (
		characters []database.Character
		name       string
		err        error
	)

	// get the first param called "names"
	names := request.URL.Query()["name"]
	if names != nil {
		name = names[0]
	}
	characters, err = database.GetCharacters(name)
	if err != nil {
		writer.Header().Set("Error", err.Error())
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	if len(characters) == 0 {
		writer.Header().Set("Error", "Character does not exist.")
		http.Error(
			writer,
			fmt.Sprintf("error. Character %s does not exist", name),
			http.StatusUnprocessableEntity,
		)
	} else {
		// encodes the characters array of structs into json and writes it to writer
		err = json.NewEncoder(writer).Encode(characters)
		if err != nil {
			writer.Header().Set("Error", err.Error())
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}
}

func HandleCharacterWithId(writer http.ResponseWriter, _ *http.Request) {
	var (
		characterWithIds []database.CharacterWithId
		err              error
	)

	characterWithIds, err = database.GetCharactersWithId()
	if err != nil {
		writer.Header().Set("Error", err.Error())
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(writer).Encode(characterWithIds)
	if err != nil {
		writer.Header().Set("Error", err.Error())
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func HandleTeapot(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusTeapot)
	_, err := writer.Write([]byte("I am a teapot"))
	if err != nil {
		// Handle the error here
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
	}
}
