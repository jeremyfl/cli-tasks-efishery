package main

import (
	"fmt"

	_ "github.com/go-kivik/couchdb" // The CouchDB driver
	"github.com/jeremylombogia/cli-tasks-efishery/internal"
)

func insertTask() {
	var task = internal.Task{"Lorem ipsum dolor si amet", false, "College"}

	var taskDone, err = internal.Store(&task)
	if err != nil {
		fmt.Println("Err")
		panic(err)
	}

	fmt.Println(taskDone)
}

func fetchTask() {
	var rows, _ = internal.Fetch()

	for rows.Next() {
		var tasks internal.Task

		if err := rows.ScanDoc(&tasks); err != nil {
			panic(err)
		}

		fmt.Println(tasks)
	}

	if rows.Err() != nil {
		panic(rows.Err())
	}
}

func main() {
	fetchTask()
}
