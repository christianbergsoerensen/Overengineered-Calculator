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

// selects postgres when using Docker (and by extension Render, DATABASE_URL is set in the Render env)
// When running locally set DB_TYPE in .env to sqlite
// Can also create a Docker container, but DATABASE_URL needs to be set to the postgres URL in the docker-compose.yml
func InitStorage() (storage.StorageInterface, error) {
	dbType := os.Getenv("DB_TYPE")

	switch dbType {
	case "postgres":
		dbURL := os.Getenv("DATABASE_URL")
		return storage.NewPostgreSQLStorage(dbURL)
	case "sqlite", "":
		return storage.NewSQLiteStorage("./internal/storage/calculations.db")
	default:
		return nil, fmt.Errorf("unsupported DB_TYPE: %s", dbType)
	}
}

func main() {
	calc := calculator.NewCalculator()

	store, err := InitStorage()
	if err != nil {
		log.Fatal("Error setting up storage: ", err)
	}

	r := api.NewRouter(calc, store)

	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "8080"
	}

	fmt.Println("Starting server on port " + portString)

	err = http.ListenAndServe(":"+portString, r)
	if err != nil {
		log.Fatal("Error setting up the server", err)
	}
}
