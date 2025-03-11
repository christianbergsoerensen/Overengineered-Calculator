package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/christianbergsoerensen/Overengineered-Calculator/internal/api"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	portString := os.Getenv("PORT")
	r := api.NewRouter()

	fmt.Println("Starting server on port " + portString)

	err = http.ListenAndServe(":"+portString, r)
	if err != nil {
		log.Fatal("Error setting up the server", err)
	}
}
