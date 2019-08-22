package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jeremylombogia/cli-tasks-efishery/cmd"
	"github.com/jeremylombogia/cli-tasks-efishery/internal"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Todo:")
	fmt.Println("---------------------")

	fmt.Println("1. List task")
	fmt.Println("2. Update task")
	fmt.Println("3. Syncing")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("1", text) == 0 {
			cmd.FetchTask()
		} else if strings.Compare("2", text) == 0 {
			fmt.Println("Update tasks")

			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.Replace(id, "\n", "", -1)

			cmd.UpdateTask(id)
		} else if strings.Compare("3", text) == 0 {
			fmt.Println("Syncing tasks")
			internal.Replicate()
		}
	}

	// cmd.InsertTask()
}
