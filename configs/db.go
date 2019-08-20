package configs

import (
	"context"
	"log"
	"os"

	_ "github.com/go-kivik/couchdb" // The CouchDB driver
	"github.com/go-kivik/kivik"
	"github.com/joho/godotenv"
)

func Init() *kivik.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := kivik.New("couch", os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}

	var db = client.DB(context.TODO(), os.Getenv("DB_NAME"))

	return db
}
