package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func getFileReference(filePath string) *os.File {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		logFile, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer logFile.Close()
	}

	fileReference, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return fileReference
}

func logData(file *os.File, heading string, content string) {
	splitString := strings.Repeat("-", len(heading))
	fmt.Fprintf(file, "%s\n%s\n%s\n%s\n\n", splitString, heading, splitString, content)
}

func runAllScript() {
	date := time.Now().Format("2006-01-02 15:04:05")
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	logFile := getFileReference(filepath.Join(pwd, "scripts.log"))
	defer logFile.Close()

	if err := os.Chdir(filepath.Join(pwd, "scripts")); err != nil {

		logData(
			logFile,
			fmt.Sprintf("[ FATAL: %s / %s ]", "folder scripts", date),
			"folder scripts not found,\nplease create folder scripts",
		)

		log.Fatal(err)
	}

	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() || strings.HasPrefix(file.Name(), "_") {
			continue
		}

		cmd := exec.Command(
			"bash", "-e",
			filepath.Join(pwd, "scripts", file.Name()),
		)
		out, err := cmd.Output()

		if err != nil {
			logData(
				logFile,
				fmt.Sprintf("[ INFO: %s / %s ]", file.Name(), date),
				strings.TrimSpace(string(err.Error())),
			)
			continue
		}

		logData(
			logFile,
			fmt.Sprintf("[ ERROR: %s / %s ]", file.Name(), date),
			strings.TrimSpace(string(out)),
		)
	}
}

func main() {
	runAllScript()
}
