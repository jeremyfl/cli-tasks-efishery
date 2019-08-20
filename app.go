package main

import (
	"context"
	"fmt"

	// "github.com/flimzy/kivik"       // Stable version of Kivik
	_ "github.com/go-kivik/couchdb" // The CouchDB driver
	"github.com/go-kivik/kivik"     // Development version of Kivik
)

func main() {
	client, err := kivik.New("couch", "http://admin:iniadmin@13.250.43.79:5984/")

	if err != nil {
		panic(err)
	}

	db := client.DB(context.TODO(), "efishery_task_test")

	doc := map[string]interface{}{
		"_id":      "cow",
		"feet":     4,
		"greeting": "moo",
	}

	rev, err := db.Put(context.TODO(), "cow", doc)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Cow inserted with revision %s\n", rev)
}
