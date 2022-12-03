package utils

import (
	"log"
	"os"
	"strings"
)

func ReadFileToStringAndSplit(filePath string, sep string) []string {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(bytes), sep)
}
