package handlers

import "net/http"

func HandleRoot(writer http.ResponseWriter, _ *http.Request) {
	EnableCors(&writer)

	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write([]byte("Welcome to the Tekken 8 API"))
	if err != nil {
		// Handle the error here
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
	}
}
