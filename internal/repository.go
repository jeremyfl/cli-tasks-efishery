package internal

import (
	"context"
	"os"

	"github.com/go-kivik/kivik"
	"github.com/jeremylombogia/cli-tasks-efishery/configs"
)

// Initialize database from config
var client, db = configs.Init()

// Fetch all data and return it as a rows
func Fetch() (*kivik.Rows, error) {
	var rows, err = db.AllDocs(context.TODO(), kivik.Options{"include_docs": true})

	return rows, err
}

// Store document and return it as a struct
func Store(task *Task) (Task, error) {
	_, _, err := db.CreateDoc(context.TODO(), task)

	return *task, err
}

func Show(id string) *kivik.Row {
	var row = db.Get(context.TODO(), id)

	return row
}

// Update document and return it as a struct document filled
func Update(id string, task *Task) (Task, error) {
	task = task
	task.IsDone = true

	_, err := db.Put(context.TODO(), id, task)

	return *task, err
}

// Replicate it replicate the target db from local db
func Replicate() {
	var _, err = client.Replicate(context.TODO(), os.Getenv("DB_URL"), os.Getenv("DB_LOCAL_URL")+os.Getenv("DB_NAME"))
	if err != nil {
		panic(err)
	}

}
