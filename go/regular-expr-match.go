package main
import "fmt"
func isMatchInit() {
	result := isMatch("aa", "a")
	fmt.Println(result)
}

func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)

	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
/*
If the current character in the pattern is '', then either the '' matches zero of the preceding element, in which case dp[i][j] = dp[i][j-2], or the '*' matches one or more of the preceding element, in which case dp[i][j] = (s[i-1] == p[j-2] || p[j-2] == '.') && dp[i-1][j].

If the current character in the pattern is not '*', then the current character in the input string and the pattern must match, i.e., s[i-1] == p[j-1] or p[j-1] == '.', and the previous characters in the input string and pattern must also match, i.e., dp[i-1][j-1].
*/
	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] || (i > 0 && (s[i-1] == p[j-2] || p[j-2] == '.') && dp[i-1][j])
			} else {
				dp[i][j] = i > 0 && (s[i-1] == p[j-1] || p[j-1] == '.') && dp[i-1][j-1]
			}
		}
	}

	return dp[m][n]
}