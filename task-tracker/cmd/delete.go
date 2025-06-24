package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"task-cli/task"
)

func DeleteTaskCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a task",
		Long: `Delete a task by ID

		Example:
		task-tracker delete 1 
		`,

		RunE: func(cmd *cobra.Command, args []string) error {
			return RunDeleteTaskCmd(args)
		},
	}
	return cmd
}

func RunDeleteTaskCmd(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("please provide a task ID")
	}

	taskID := args[0]
	taskIDInt, err := strconv.ParseInt(taskID, 10, 32)
	if err != nil {
		return err
	}
	return task.DeleteTask(taskIDInt)
}
