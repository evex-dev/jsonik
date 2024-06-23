package runner

import (
	"encoding/json"
	"fmt"
	"jsonik/logger"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
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

	fmt.Println(logger.InfoStyle.Render(logger.InfoMark, "Searching tasks :", pickTaskNameFromFilePath(taskName)))
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
		fmt.Println(logger.ErrorStyle.Render(logger.ErrorMark, "Tasks not found :", pickTaskNameFromFilePath(taskName)))
		return
	}

	if len(files) > 1 {
		fmt.Println(logger.ErrorStyle.Render(logger.ErrorMark, "Multiple tasks found :", pickTaskNameFromFilePath(taskName)))
		return
	}

	fmt.Println(logger.SuccessStyle.Render(logger.SuccessMark, "Found tasks :", pickTaskNameFromFilePath(taskName)))

	file, err := os.ReadFile(files[0])

	if err != nil {
		fmt.Println(logger.ErrorStyle.Render(logger.ErrorMark, err.Error()))
		return
	}

	var TaskListJSON TaskList
	json.Unmarshal(file, &TaskListJSON)
	fmt.Println()
	fmt.Println(logger.LoadingStyle.Render(logger.LoadingMark, "Running tasks :", pickTaskNameFromFilePath(taskName)))

	taskListWaitlist := TaskListJSON.Tasks
	taskListAligned := []Task{}

	for {
		for _, task := range taskListWaitlist {
			isHaveNeeds := task.Needs != nil

			if isHaveNeeds {
				needsTaskList := task.Needs
				endedNeeds := 0
				needsEndedNeeds := len(needsTaskList)

				for _, taskAligned := range taskListAligned {
					if slices.Contains(needsTaskList, taskAligned.Label) {
						endedNeeds++
					}

					if endedNeeds == needsEndedNeeds {
						break
					}
				}

				if endedNeeds == needsEndedNeeds {
					taskListWaitlist = deleteValueFromSlice(taskListWaitlist, task)
					taskListAligned = append(taskListAligned, task)
				} else {
					taskListWaitlist = deleteValueFromSlice(taskListWaitlist, task)
					taskListWaitlist = append(taskListWaitlist, task)
				}
			} else {
				taskListWaitlist = deleteValueFromSlice(taskListWaitlist, task)
				taskListAligned = append(taskListAligned, task)
			}
		}

		if len(taskListWaitlist) == 0 {
			break
		}
	}

	for _, task := range taskListAligned {
		fmt.Println()
		fmt.Println(logger.LoadingStyle.Render(logger.LoadingMark, "Running task :", task.Label))

		for _, command := range task.Run {
			fmt.Println(logger.LoadingStyle.Render("", logger.GraftingMart, logger.LoadingMark, "Running command :", command))
			
			parsedCommand := parseCommand(command)
			cmd := exec.Command(parsedCommand[0], parsedCommand[1:]...)

			if err := cmd.Run(); err != nil {
				fmt.Println("", logger.ErrorStyle.Render(logger.GraftingMart), logger.ErrorStyle.Render(logger.ErrorMark, err.Error()))
			}
		}
	}
}

func isPath(taskName string) bool {
	return strings.HasSuffix(taskName, ".jsonik.json") || strings.HasSuffix(taskName, ".jk.json")
}

func pickTaskNameFromFilePath(filePath string) string {
	splitedFilePath := strings.Split(strings.ReplaceAll(filePath, string(filepath.Separator), "/"), "/")
	fileName := splitedFilePath[len(splitedFilePath)-1]
	return strings.ReplaceAll(strings.ReplaceAll(fileName, ".jk.json", ""), ".jsonik.json", "")
}

func deleteValueFromSlice(slice []Task, value Task) []Task {
	slice2 := []Task{}
	for i, v := range slice {
		if v.Label != value.Label {
			slice2 = append(slice2, slice[i])
		}
	}
	return slice2
}

func parseCommand(command string) []string {
	splitedCommand := strings.Split(command, " ")
	return []string{splitedCommand[0]}
}
