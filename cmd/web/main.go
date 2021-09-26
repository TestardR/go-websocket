package main

import (
	"log"
	"net/http"

	"github.com/TestardR/go-websocket/internal/handlers"
)

func main() {
	mux := routes()

	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()

	log.Println("Starting web server on port 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println(err)
	}
}
