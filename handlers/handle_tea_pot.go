package handlers

import "net/http"

func HandleTeapot(writer http.ResponseWriter, _ *http.Request) {
	EnableCors(&writer)

	writer.WriteHeader(http.StatusTeapot)
	_, err := writer.Write([]byte("I am a teapot"))
	if err != nil {
		// Handle the error here
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
	}
}
func HandleRoot(writer http.ResponseWriter, _ *http.Request) {
	EnableCors(&writer)

	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write([]byte("Hello, World"))
	if err != nil {
		// Handle the error here
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
	}
}
