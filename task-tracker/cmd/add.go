package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"task-cli/task"
)

func AddNewTaskCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a task to the task list",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunAddTaskCmd(args)
		},
	}
	return cmd
}

func RunAddTaskCmd(args []string) error {
	if len(args) == 0 {
		return errors.New("task description is required")
	}

	description := args[0]
	id, err := task.GetNewTaskID()
	if err != nil {
		return errors.New(err.Error())
	}
	newTask := task.NewTask(id, description)
	err = task.AddNewTask(newTask)
	if err != nil {
		return err
	}
	fmt.Printf("Task added successfully (ID: %d)\n", id)
	return nil
}
