package runner

import (
	"encoding/json"
	"fmt"
	"jsonik/logger"
	"os"
	"path/filepath"
	"strings"
)

type Task struct {
	Label string   `json:"label"`
	Needs []string `json:"needs"`
	Run   []string `json:"run"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

func Runner(taskName string) {
	if strings.Trim(taskName, " ") == "" || strings.Trim(taskName, " ") == "." {
		taskName = "main"
	} else {
		taskName = strings.Trim(taskName, " ")
	}

	fmt.Println(logger.InfoStyle.Render(logger.InfoMark, "Searching task :", pickTaskNameFromFilePath(taskName)))
	patterns := []string{}

	if isPath(taskName) {
		patterns = []string{taskName}
	} else {
		patterns = []string{
			"./" + taskName + ".jsonik.json",
			"./" + taskName + ".jk.json",
			"./**/" + taskName + ".jsonik.json",
			"./**/" + taskName + ".jk.json",
		}
	}

	files := []string{}

	for _, pattern := range patterns {
		foundFiles, err := filepath.Glob(pattern)
		if err != nil {
			fmt.Println(logger.ErrorStyle.Render(logger.ErrorMark, err.Error()))
			return
		}
		files = append(files, foundFiles...)
	}

	if len(files) == 0 {
		fmt.Println(logger.ErrorStyle.Render(logger.ErrorMark, "Task not found :", pickTaskNameFromFilePath(taskName)))
		return
	}

	if len(files) > 1 {
		fmt.Println(logger.ErrorStyle.Render(logger.ErrorMark, "Multiple tasks found :", pickTaskNameFromFilePath(taskName)))
		return
	}

	fmt.Println(logger.SuccessStyle.Render(logger.SuccessMark, "Found task :", pickTaskNameFromFilePath(taskName)))

	file, err := os.ReadFile(files[0])

	if err != nil {
		fmt.Println(logger.ErrorStyle.Render(logger.ErrorMark, err.Error()))
		return
	}

	var TaskListJSON TaskList
	json.Unmarshal(file, &TaskListJSON)
	fmt.Println(TaskListJSON.Tasks)
}

func isPath(taskName string) bool {
	return strings.HasSuffix(taskName, ".jsonik.json") || strings.HasSuffix(taskName, ".jk.json")
}

func pickTaskNameFromFilePath(filePath string) string {
	splitedFilePath := strings.Split(strings.ReplaceAll(filePath, string(filepath.Separator), "/"), "/")
	fileName := splitedFilePath[len(splitedFilePath)-1]
	return strings.ReplaceAll(strings.ReplaceAll(fileName, ".jk.json", ""), ".jsonik.json", "")
}
