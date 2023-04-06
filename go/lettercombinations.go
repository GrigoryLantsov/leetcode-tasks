package main
import (
	"fmt"
)
func lettercombinationinit() {
	result := letterCombinations("23")
	fmt.Println(result)
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	// Define the mapping of digits to letters
	letters := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	// Initialize the result slice with the first set of letters
	result := make([]string, 0)
	for _, c := range letters[digits[0]-'0'] {
		result = append(result, string(c))
	}

	// Iterate over the remaining digits
	for i := 1; i < len(digits); i++ {
		newResult := make([]string, 0)
		for _, c := range letters[digits[i]-'0'] {
			for _, s := range result {
				newResult = append(newResult, s+string(c))
			}
		}
		result = newResult
	}

	return result
}