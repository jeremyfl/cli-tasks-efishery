package main

import (
	"github.com/jeremylombogia/cli-tasks-efishery/cmd"
	"github.com/jeremylombogia/cli-tasks-efishery/internal"
)

func main() {
	// cmd.InsertTask()
	// cmd.FetchTask()
	cmd.UpdateTask("bles1osllhcklvh3pm1g")

	internal.Replicate()
}
