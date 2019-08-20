package cmd

import (
	"fmt"

	"github.com/jeremylombogia/cli-tasks-efishery/internal"
)

func InsertTask() {
	var task = internal.Task{"Lorem ipsum dolor si amet", false, "College"}

	var taskDone, err = internal.Store(&task)
	if err != nil {
		fmt.Println("Err")
		panic(err)
	}

	fmt.Println(taskDone)
}

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
