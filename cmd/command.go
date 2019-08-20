package cmd

import (
	"fmt"

	"github.com/jeremylombogia/cli-tasks-efishery/internal"
)

// InsertTask it store repositories and return a struct model of it task
func InsertTask() {
	var task = internal.Task{"Lorem ipsum dolor si amet", false, "College"}

	var taskDone, err = internal.Store(&task)
	if err != nil {
		panic(err)
	}

	fmt.Println(taskDone)
}

// FetchTask read all of document as a struct model
func FetchTask() {
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
