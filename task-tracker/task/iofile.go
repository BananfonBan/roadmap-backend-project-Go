/*
input/output file
*/
package task

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
)

func GetFileTaskPath() string {
	return getFileTaskPath()
}

func getFileTaskPath() string {
	cwd, error := os.Getwd()
	if error != nil {
		fmt.Println("Error getting current working directory:", error)
		return ""
	}
	return path.Join(cwd, "tasks.json")
}

func WriteTasksToFile(task []Task) error {
	filePath := getFileTaskPath()
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	err = json.NewEncoder(file).Encode(&task)
	if err != nil {
		fmt.Println("Error encoding:", err)
		return err
	}
	return nil
}

func ReadTaskFromFile() ([]Task, error) {
	filePath := getFileTaskPath()
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println("File does not exist. Creating file...")
		file, err := os.Create(filePath)
		os.WriteFile(filePath, []byte(""), os.ModeAppend.Perm())

		if err != nil {
			fmt.Println("Error creating file:", err)
			return nil, err
		}

		defer file.Close()

		return []Task{}, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()
	tasks := []Task{}
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil && err != io.EOF {
		fmt.Println("Error decoding file:", err)
		return nil, err
	}
	return tasks, nil
}
