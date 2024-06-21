package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"pns/cmd/pns/router"
)

const (
	defaultPort = "8080"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading environment file: ", err)
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Println("No port set, using default one (" + defaultPort + ")")
		port = defaultPort
	}

	log.Printf("Listening on port %s...", port)
	err := http.ListenAndServe(":"+port, router.New())
	if err != nil {
		log.Fatal(err)
	}
}
