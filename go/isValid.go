package main
import (
	"fmt"
)

func isValidInit() {
	s := "(]"
	isValid := isValid(s)
	fmt.Println(isValid)
}


func isValid(s string) bool {
	// create a stack to keep track of the opening parentheses
	stack := make([]rune, 0)

	// iterate through each character in the string
	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			// if the character is an opening parenthesis, push it onto the stack
			stack = append(stack, char)
		} else if char == ')' || char == '}' || char == ']' {
			// if the character is a closing parenthesis, pop the last opening parenthesis from the stack
			if len(stack) == 0 {
				// if the stack is empty, there is no corresponding opening parenthesis, so the string is not valid
				return false
			}
			lastOpen := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if char == ')' && lastOpen != '(' || char == '}' && lastOpen != '{' || char == ']' && lastOpen != '[' {
				// if the closing parenthesis does not match the last opening parenthesis, the string is not valid
				return false
			}
		}
	}

	// if the stack is not empty, there are opening parentheses with no corresponding closing parentheses, so the string is not valid
	return len(stack) == 0
}