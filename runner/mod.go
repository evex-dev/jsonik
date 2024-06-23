package runner

import (
	"fmt"
	"jsonik/logger"
	"path/filepath"
	"strings"
)

func Runner(taskName string) {
	fmt.Println(logger.InfoStyle.Render(logger.InfoMark, "Searching task :", taskName))
	patterns := []string{}

	if isPath(taskName) {
		patterns = []string{taskName}
	}else {
		patterns = []string{
			"./" + taskName + ".jsonik.json",
			"./" + taskName + ".jk.json",
			"./**/" + taskName + ".jsonik.json",
			"./**/" + taskName + ".jk.json",
		}
	}

	files := []string{}

	for _, pattern := range patterns {
		foundFiles, err := findFiles(pattern)
		if err != nil {
			fmt.Println(logger.ErrorStyle.Render(logger.ErrorMark, err.Error()))
			return
		}
		files = append(files, foundFiles...)
	}
	
	fmt.Println(patterns, files)
}

func isPath(taskName string) bool {
	if strings.HasSuffix(taskName, ".jsonik.json") || strings.HasSuffix(taskName, ".jk.json") {
		return true
	}
	return false
}

func findFiles(pattern string) ([]string, error) {
	files, err := filepath.Glob(pattern)
	return files, err
}
