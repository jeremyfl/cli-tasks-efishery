package configs

import (
	"context"
	"log"
	"os"

	_ "github.com/go-kivik/couchdb"
	"github.com/go-kivik/kivik"
	"github.com/joho/godotenv"
)

// Init the database and return as kivik.DB
func Init() (*kivik.Client, *kivik.DB) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := kivik.New("couch", os.Getenv("DB_LOCAL_URL"))
	if err != nil {
		panic(err)
	}

	var db = client.DB(context.TODO(), os.Getenv("DB_NAME"))

	return client, db
}
