package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"task-cli/task"
)

func UpdateTaskDescriptionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a task description",
		Long: `Update a task description by ID

		
		Example: task-cli update 1 "New description"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateTaskDescriptionCmd(args)
		},
	}
	return cmd
}

func RunUpdateTaskDescriptionCmd(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("please provide a task ID and the new description")
	}
	taskID := args[0]
	taskIDInt, err := strconv.ParseInt(taskID, 10, 32)
	if err != nil {
		return err
	}
	newDescription := args[1]
	return task.UpdateTaskDescription(taskIDInt, newDescription)
}

func UpdateTaskStatusTodoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-todo",
		Short: "Mark a task as todo",
		Long: `Update a task status by ID on "todo"

		Example: task-cli mark-todo 1`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateTaskStatusCmd(args, task.TASK_STATUS_TODO)
		},
	}
	return cmd
}

func UpdateTaskStatusInProgressCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-in-progress",
		Short: "Mark a task as in-progress",
		Long: `Update a task status by ID on "in-progress"
		
		Example: task-cli mark-in-progress 1`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateTaskStatusCmd(args, task.TASK_STATUS_IN_PROGRESS)
		},
	}
	return cmd
}

func UpdateTaskStatusDoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-done",
		Short: "Mark a task as done",
		Long: `Update a task status by ID on "done"
		
		Example: task-cli mark-done 1`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateTaskStatusCmd(args, task.TASK_STATUS_DONE)
		},
	}
	return cmd
}

func RunUpdateTaskStatusCmd(args []string, status task.TaskStatus) error {
	if len(args) == 0 {
		return fmt.Errorf("task ID is required")
	}
	taskID := args[0]
	taskIDInt, err := strconv.ParseInt(taskID, 10, 32)
	if err != nil {
		return err
	}
	return task.UpdateTaskStatus(taskIDInt, status)
}
