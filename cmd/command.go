package cmd

import (
	"fmt"

	"github.com/jeremylombogia/cli-tasks-efishery/internal"
	"github.com/rs/xid"
)

// InsertTask it store repositories and return a struct model of it task
func InsertTask() {
	guid := xid.New()

	var task = internal.Task{
		ID:      guid.String(),
		Content: "Lorem ipsum dolor si amet",
		IsDone:  false,
		Tag:     "College",
	}

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

		fmt.Print("Content: ")
		fmt.Println(tasks.Content)
		fmt.Print("Done?: ")
		fmt.Println(tasks.IsDone)
		fmt.Print("Tag: ")
		fmt.Println(tasks.Tag)
		fmt.Println()
	}

	if rows.Err() != nil {
		panic(rows.Err())
	}
}

// UpdateTask from command into done
func UpdateTask(id string) {
	var task internal.Task

	var row = internal.Show(id)

	if err := row.ScanDoc(&task); err != nil {
		panic(err)
	}

	var _, err = internal.Update(id, &task)
	if err != nil {
		panic(err)
	}

	fmt.Println(task)
}
