package utils

import (
	"log"
	"strconv"
)

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

func IntersectBetweenStrings(first string, second string) (keys []rune) {
	set := make(map[rune]int)
	hash := make(map[rune]int)
	for _, v := range first {
		hash[v] += 1
	}

	for _, v := range second {
		if _, ok := hash[v]; ok {
			set[v] += 1
		}
	}
	for k := range set {
		keys = append(keys, k)
	}
	return keys
}

func SplitStringByIndex(input string, start int, end int) string {
	return input[start:end]
}
