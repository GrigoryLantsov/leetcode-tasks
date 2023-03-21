package main
import "fmt"
func romanint() {
	result := romanToInt("VVVVIII")
	fmt.Println(result)
}
func romanToInt(s string) int {
	// create a map to store the values of each Roman numeral
	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
		}

		result := 0
		prev := 0

		// iterate through the string in reverse order
		for i := len(s) - 1; i >= 0; i-- {
			curr := values[s[i]]
			if curr < prev {
				result -= curr
			} else {
				result += curr
			}
			prev = curr
		}

		return result
}