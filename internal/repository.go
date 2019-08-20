package internal

import (
	"context"

	"github.com/go-kivik/kivik"
	"github.com/jeremylombogia/cli-tasks-efishery/configs"
)

var db = configs.Init()

// Fetch all data and return it rows
func Fetch() (*kivik.Rows, error) {
	rows, err := db.AllDocs(context.TODO(), kivik.Options{"include_docs": true})

	return rows, err
}

// Store database with struct
func Store(task *Task) (Task, error) {
	_, _, err := db.CreateDoc(context.TODO(), task)

	return *task, err
}

func Update() {

}
