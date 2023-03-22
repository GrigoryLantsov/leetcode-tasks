package main
import "fmt"
func lengthsubstringinit() {
	result := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(result)
}

func lengthOfLongestSubstring(s string) int {
	start := 0
	end := 0
	maxLen := 0
	seen := make(map[byte]bool)

	for end < len(s) {
		if !seen[s[end]] {
			seen[s[end]] = true
			end++
			maxLen = max(maxLen, end-start)
		} else {
			delete(seen, s[start])
			start++
		}
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}