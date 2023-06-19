package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	var combinations []string
	backtrack(&combinations, "", 0, 0, n)
	return combinations
}

func backtrack(combinations *[]string, currentCombination string, open, close, max int) {
	if len(currentCombination) == max*2 {
		*combinations = append(*combinations, currentCombination)
		return
	}

	if open < max {
		backtrack(combinations, currentCombination+"(", open+1, close, max)
	}
	if close < open {
		backtrack(combinations, currentCombination+")", open, close+1, max)
	}
}

func generateParenthesisInit() {
	fmt.Println(generateParenthesis(3))  // Output: ["((()))","(()())","(())()","()(())","()()()"]
	fmt.Println(generateParenthesis(1))  // Output: ["()"]
}
