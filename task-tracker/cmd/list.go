package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"task-cli/task"
)

func ListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Long: `List all tasks. You can filter tasks by status

    Example:
    task-tracker list todo
    task-tracker list in-progress
    task-tracker list done
    `,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunListTaskCmd(args)
		},
	}
	return cmd
}

func RunListTaskCmd(args []string) error {
	if len(args) == 0 {
		alltasks, err := task.ReadTaskFromFile()
		if err != nil {
			return err
		}
		return printListTasks(alltasks)
	}
	status := task.TaskStatus(args[0])
	switch status {
	case task.TASK_STATUS_TODO, task.TASK_STATUS_IN_PROGRESS, task.TASK_STATUS_DONE:
		goto filter
	default:
		return fmt.Errorf("the specified status does not exist")
	}
filter:
	tasks, err := task.FilterListTasks(status)
	if err != nil {
		return err
	}
	return printListTasks(tasks)
}

func printListTasks(tasks []task.Task) error {
	if len(tasks) == 0 {
		fmt.Println("Task list is empty!")
		return nil
	}
	for _, task := range tasks {
		fmt.Println()
		fmt.Printf("ID: %d\n", task.ID)
		fmt.Printf("Description: %s\n", task.Description)
		fmt.Printf("Status: %s\n", task.Status)
		fmt.Printf("Created at: %s\n", task.CreatedAt)
		fmt.Printf("Updated at: %s\n", task.UpdatedAt)
	}
	fmt.Println()
	return nil
}
