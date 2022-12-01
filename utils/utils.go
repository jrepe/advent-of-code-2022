package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFileToStringAndSplit(filePath string, sep string) []string {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(bytes), sep)
}

func ConvertStringSliceToInt(in []string) []int {
	out := make([]int, len(in))
	for _, i := range in {
		num, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		out = append(out, num)
	}
	return out
}

func ConvertStringSliceToInt64(in []string) []int64 {
	out := make([]int64, len(in))
	for _, i := range in {
		num, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		out = append(out, num)
	}
	return out
}
