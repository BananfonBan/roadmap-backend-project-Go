package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "task-cli",
		Short: "Simple task-tracker for CLI",
		Long:  `Task Tracker is a CLI tool for managing tasks. It allows you to create, list, delete and update status and description tasks.`,
	}
	cmd.AddCommand(AddNewTaskCmd())
	cmd.AddCommand(DeleteTaskCmd())
	cmd.AddCommand(UpdateTaskDescriptionCmd())
	cmd.AddCommand(UpdateTaskStatusTodoCmd())
	cmd.AddCommand(UpdateTaskStatusInProgressCmd())
	cmd.AddCommand(UpdateTaskStatusDoneCmd())
	cmd.AddCommand(ListCmd())
	return cmd
}
