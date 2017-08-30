package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	origins := handlers.AllowedOrigins(AllowedOrigins)
	router := NewRouter()
	router.Host(BaseURL)
	log.Fatal(http.ListenAndServe(":8080",
		handlers.CORS(origins)(router)))
}
