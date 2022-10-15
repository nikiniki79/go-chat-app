package main

import (
	"go-chat-app/handlers"
	"log"
	"net/http"
)

func main() {
	mux := routes()

	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()

	log.Println("Starting web server on port 9080")

	_ = http.ListenAndServe(":9080", mux)
}
