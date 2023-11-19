package handlers

import "net/http"

// EnableCors enables cors from the frontend host
func EnableCors(writer *http.ResponseWriter) {
	(*writer).Header().Set("Access-Control-Allow-Origin", "*")
}
