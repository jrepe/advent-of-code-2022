package utils

func SumOfIntItems(items []int) int {
	sum := 0
	for _, i := range items {
		sum += i
	}
	return sum
}
