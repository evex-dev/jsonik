package main

import (
	"fmt"
	"jsonik/logger"
	"jsonik/runner"
	"os"
)

func main() {
	args := os.Args[1:]
	commandName := "unknown"

	if len(args) == 2 && args[0] == "run" {
		commandName = "run"
	} else if len(args) == 1 && args[0] == "help" {
		commandName = "help"
	}

	switch commandName {
	case "run":
		taskName := args[1]
		runner.Runner(taskName)
	case "help":
		fmt.Println(logger.InfoStyle.Render(logger.InfoMark, "Use 'jsonik run {task_title}' to run a task"))
		fmt.Println(logger.InfoStyle.Render(logger.InfoMark, "Use 'jsonik help' for help"))
	default:
		fmt.Println(logger.ErrorStyle.Render(logger.ErrorMark, "Unknown command : ", commandName))
		fmt.Println(logger.InfoStyle.Render(logger.InfoMark, "Use 'jsonik help' for help"))
	}
}
