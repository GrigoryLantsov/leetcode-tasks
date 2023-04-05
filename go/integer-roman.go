package main
import "fmt"
func integerromanint() {
	result := intToRoman(3)
	fmt.Println(result)
}
func intToRoman(num int) string {
	// create two slices for the Roman numerals and their corresponding values
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	// initialize result variable
	roman := ""

	// iterate over the values and symbols, adding the symbols to the result while subtracting the corresponding value from the input
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			roman += symbols[i]
			num -= values[i]
		}
	}

	// return the result
	return roman
}