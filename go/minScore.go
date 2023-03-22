package main
import (
	"fmt"
	"math"
)

func minScoreinit() {
	n := 4
	roads := [][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}}
	score := minScore(n, roads)
	fmt.Println(score)
}

func minScore(n int, roads [][]int) int {
	dest := make([]int, n)
	var union func(int, int)
	var search func(int) int
	union = func(x, y int) {
		dest[search(x)] = search(y)
	}
	search = func(x int) int {
		if dest[x] != x {
			dest[x] = search(dest[x])
		}
		return dest[x]
	}

	for i := 0; i < n; i++ {
		dest[i] = i
	}

	for _, r := range roads {
		union(r[0]-1, r[1]-1)
	}

	t := search(0)
	result := math.MaxInt
	for _, r := range roads {
		if search(r[0]-1) != t || search(r[1]-1) != t {
			continue
		}
		result = min(result, r[2])
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}