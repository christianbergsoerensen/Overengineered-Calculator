package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/christianbergsoerensen/Overengineered-Calculator/internal/api"
	"github.com/christianbergsoerensen/Overengineered-Calculator/internal/calculator"
	"github.com/christianbergsoerensen/Overengineered-Calculator/internal/storage"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	calc := calculator.NewCalculator()
	store, err := storage.NewSQLiteStorage("calculations.db")
	if err != nil {
		log.Fatal("Error setting up the storage ", err)
	}
	r := api.NewRouter(calc, store)

	portString := os.Getenv("PORT")
	fmt.Println("Starting server on port " + portString)

	err = http.ListenAndServe(":"+portString, r)
	if err != nil {
		log.Fatal("Error setting up the server", err)
	}
}
