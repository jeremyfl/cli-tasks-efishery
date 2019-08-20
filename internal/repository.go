package internal

import (
	"context"

	"github.com/go-kivik/kivik"
	"github.com/jeremylombogia/cli-tasks-efishery/configs"
)

var db = configs.Init()

func Fetch() (*kivik.Rows, error) {
	rows, err := db.AllDocs(context.TODO(), kivik.Options{"include_docs": true})

	if err != nil {
		panic(err)
	}

	return rows, err
}

// Store database with struct
func Store(task *Task) (Task, error) {
	_, _, err := db.CreateDoc(context.TODO(), task)

	if err != nil {
		panic(err)
	}

	return *task, err
}

func Update() {

}
