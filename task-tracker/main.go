package main

import (
	"fmt"
	"task-cli/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
