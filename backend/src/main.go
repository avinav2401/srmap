package main

import (
	"log"
	"net"
	"os"

	"goscraper/src/bootstrap"
)

func main() {
	app := bootstrap.SetupApp()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s...", port)
	ln, err := net.Listen("tcp", "[::]:"+port)
	if err != nil {
		log.Fatalf("Failed to bind: %v", err)
	}
	log.Printf("Starting server on port %s...", port)
	if err := app.Listener(ln); err != nil {
		log.Printf("Server error: %+v", err)
	}
}
