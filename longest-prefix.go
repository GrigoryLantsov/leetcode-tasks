package main
import "fmt"
func prefixinit() {
	result := longestCommonPrefix([]string{"flower","flow","flight"})
	fmt.Println(result)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		// if the prefix is empty, there is no common prefix
		if prefix == "" {
			return ""
		}

		// find the index of the first character that is different
		j := 0
		for ; j < len(prefix) && j < len(strs[i]); j++ {
			if prefix[j] != strs[i][j] {
				break
			}
		}

		// update the prefix to the common prefix
		prefix = prefix[:j]
	}

	return prefix
}