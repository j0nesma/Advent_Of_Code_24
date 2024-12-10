package FileReader

import (
	"fmt"
	"os"
	"strings"

	"Advent_Of_Code_24/Error"
)

func ReadFileToString(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	Error.Check(err)
	return string(content), nil
}

func ReadFileToArray(filePath string) []string {
	toString, err := ReadFileToString(filePath)
	fmt.Println(toString)
	Error.Check(err)
	return strings.Split(toString, "\n")
}
