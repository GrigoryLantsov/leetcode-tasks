package main
import "fmt"
func palindrome() {
	result := isPalindrome(-12221)
	fmt.Println(result)
}
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reversed := 0
	original := x
	for x > 0 {
		digit := x % 10
		reversed = reversed*10 + digit
		x /= 10
	}
	return reversed == original
}
