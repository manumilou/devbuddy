package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	// "github.com/pior/dad/pkg/config"
	"github.com/pior/dad/pkg/project"
	"github.com/pior/dad/pkg/tasks"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Ensure the project is up and running",
	Run:   upRun,
	// Args:  OnlyOneArg,
}

func upRun(cmd *cobra.Command, args []string) {
	// conf := config.Load()

	path, err := os.Getwd()
	checkError(err)

	proj, err := project.FindCurrent(path)
	checkError(err)

	var taskList []tasks.Task

	for _, taskdef := range proj.Manifest.Up {
		// fmt.Printf("taskdef: %+v\n", taskdef)
		task, err := tasks.BuildFromDefinition(taskdef)
		checkError(err)
		taskList = append(taskList, task)
		// fmt.Printf("task: %+v\n", task)
	}

	for _, task := range taskList {
		fmt.Printf("Running task: %+v\n", task)
		task.Perform()
	}

}