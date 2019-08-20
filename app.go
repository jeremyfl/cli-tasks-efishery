package main

import (
	"context"
	"fmt"

	_ "github.com/go-kivik/couchdb" // The CouchDB driver
	"github.com/go-kivik/kivik"
	"github.com/jeremylombogia/cli-tasks-efishery/configs"
	"github.com/jeremylombogia/cli-tasks-efishery/internal"
)

func main() {
	db := configs.Init()

	task := internal.Task{
		Content: "Lorem ipsum",
		IsDone:  false,
	}

	_, _, err := db.CreateDoc(context.TODO(), task)
	if err != nil {
		panic(err)
	}

	rows, err := db.AllDocs(context.TODO(), kivik.Options{"include_docs": true})

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var tasks internal.Task

		if err = rows.ScanDoc(&tasks); err != nil {
			panic(err)
		}

		fmt.Println(tasks)
	}

	if rows.Err() != nil {
		panic(rows.Err())
	}
}
