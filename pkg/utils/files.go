package utils

import "os"

func getFileReference(filePath string) *os.File {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		logFile, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}
		defer logFile.Close()
	}

	fileReference, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	return fileReference
}
