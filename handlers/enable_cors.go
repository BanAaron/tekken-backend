package handlers

import "net/http"

// frontendHost should be the IP Address and port of the front end app
var frontendHost = "http://localhost:5173"

// EnableCors enables cors from the frontend host
func EnableCors(writer *http.ResponseWriter) {
	(*writer).Header().Set("Access-Control-Allow-Origin", frontendHost)
}
