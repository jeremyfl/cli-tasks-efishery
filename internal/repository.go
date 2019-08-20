package internal

import (
	"context"

	"github.com/go-kivik/kivik"
	"github.com/jeremylombogia/cli-tasks-efishery/configs"
)

// Initialize database from config
var db = configs.Init()

// Fetch all data and return it as a rows
func Fetch() (*kivik.Rows, error) {
	rows, err := db.AllDocs(context.TODO(), kivik.Options{"include_docs": true})

	return rows, err
}

// Store document and return it as a struct
func Store(task *Task) (Task, error) {
	_, _, err := db.CreateDoc(context.TODO(), task)

	return *task, err
}

// Update document and return it as a struct document filled
func Update() {

}
