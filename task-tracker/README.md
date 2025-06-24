# Task-tracker

Solution for project [task-tracker from roadmap.sh](https://roadmap.sh/projects/task-tracker)

Simple CLI utility for tracker your task

## Features

Add, delete and update status and description tasks

## Build project

```
cd task-tracker
go build -o task-tracker
```

This command create a executable file task-tracker

```
./task-tracker help
Task Tracker is a CLI tool for managing tasks. It allows you to create, list, delete and update status and description tasks.

Usage:
  task-cli [command]

Available Commands:
  add              Add a task to the task list
  completion       Generate the autocompletion script for the specified shell
  delete           Delete a task
  help             Help about any command
  list             List all tasks
  mark-done        Mark a task as done
  mark-in-progress Mark a task as in-progress
  mark-todo        Mark a task as todo
  update           Update a task description

Flags:
  -h, --help   help for task-cli

Use "task-cli [command] --help" for more information about a command.
```
