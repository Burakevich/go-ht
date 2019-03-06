package main

// Specify MapTo function here

func MapTo(items []int, fn func(int, int) string) []string {
	var result []string
	for i, v := range items {
		result = append(result, fn(v, i))
	}
	return result
}

func Convert(arr []int) []string {
	var result []string
	for _, v := range arr {
		result = append(result, convertSingle(v))
	}
	return result
}

func convertSingle(i int) string {
	numbers := []string{"unknoun", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	if i < 1 || i > 9 {
		return "unknown"
	}
	return numbers[i]
}

func main() {
}
