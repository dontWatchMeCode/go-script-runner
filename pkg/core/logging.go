package core

import (
	"fmt"
	"os"
	"strings"
)

func LogData(file *os.File, heading string, content string) {
	splitString := strings.Repeat("-", len(heading))
	fmt.Fprintf(file, "%s\n%s\n%s\n%s\n\n", splitString, heading, splitString, content)
}
