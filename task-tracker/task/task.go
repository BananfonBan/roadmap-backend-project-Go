package task

import (
	"fmt"
	"time"
)

type TaskStatus string

const (
	TASK_STATUS_TODO        TaskStatus = "todo"
	TASK_STATUS_IN_PROGRESS TaskStatus = "in-progress"
	TASK_STATUS_DONE        TaskStatus = "done"
)

type Task struct {
	ID          int64      `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

func NewTask(id int64, description string) *Task {
	return &Task{
		ID:          id,
		Description: description,
		Status:      TASK_STATUS_TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func GetNewTaskID() (int64, error) {
	tasks, err := ReadTaskFromFile()
	if err != nil {
		return 0, err
	}

	var currentID int64
	if len(tasks) > 0 {
		lastTask := tasks[len(tasks)-1]
		currentID = lastTask.ID + 1
	} else {
		currentID = 1
	}
	return currentID, nil
}

func AddNewTask(task *Task) error {
	tasks, err := ReadTaskFromFile()
	if err != nil {
		return err
	}
	tasks = append(tasks, *task)
	return WriteTasksToFile(tasks)
}

func DeleteTask(id int64) error {
	tasks, err := ReadTaskFromFile()
	if err != nil {
		return err
	}
	var taskExist bool = false
	var updetedTasks []Task
	for i, task := range tasks {
		if task.ID == id {
			updetedTasks = append(tasks[:i], tasks[i+1:]...)
			taskExist = true
		}
	}
	if !taskExist {
		return fmt.Errorf("Task not found (ID: %d)", id)
	}
	return WriteTasksToFile(updetedTasks)
}

func UpdateTaskStatus(id int64, stasus TaskStatus) error {
	tasks, err := ReadTaskFromFile()
	if err != nil {
		return err
	}
	var taskExist bool = false
	var updetedTasks []Task
	for _, task := range tasks {
		if task.ID == id {
			taskExist = true
			task.UpdatedAt = time.Now()
			task.Status = stasus
		}
		updetedTasks = append(updetedTasks, task)
	}
	if !taskExist {
		return fmt.Errorf("Task not found (ID: %d)", id)
	}
	return WriteTasksToFile(updetedTasks)
}

func UpdateTaskDescription(id int64, description string) error {
	tasks, err := ReadTaskFromFile()
	if err != nil {
		return err
	}
	var taskExist bool = false
	var updetedTasks []Task
	for _, task := range tasks {
		if task.ID == id {
			taskExist = true
			task.UpdatedAt = time.Now()
			task.Description = description
		}
		updetedTasks = append(updetedTasks, task)
	}
	if !taskExist {
		return fmt.Errorf("Task not found (ID: %d)", id)
	}
	return WriteTasksToFile(updetedTasks)
}

func FilterListTasks(status TaskStatus) ([]Task, error) {
	tasks, err := ReadTaskFromFile()
	if err != nil {
		return []Task{}, err
	}
	filteredTasks := []Task{}
	for _, task := range tasks {
		if task.Status == status {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks, nil
}
