package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"pns/handler"
)

const (
	defaultPort    = "8080"
	tokensEndpoint = "/api/v1/tokens"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading environment file: ", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("No port set, using default one (" + defaultPort + ")")
		port = defaultPort
	}

	// setup http handlers for our API routes
	http.HandleFunc(tokensEndpoint, handler.RegisterDeviceToken)
	http.HandleFunc(tokensEndpoint, handler.DeleteDeviceToken)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
