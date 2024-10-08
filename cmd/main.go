package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gerardo02/practicando-mas/cmd/api"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("no db auth")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("cant connect to database")
	}

	apiServer := api.NewApiServer(":8080", conn)
	if err := apiServer.Run(); err != nil {
		log.Fatal(err)
	}

	log.Println("Listening on port 8081")
}
