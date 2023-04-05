package main
import "fmt"
func maxAreaInit() {
	height1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("Input:", height1)
	fmt.Println("Output:", maxArea(height1)) // expected output: 49

	height2 := []int{1, 1}
	fmt.Println("Input:", height2)
	fmt.Println("Output:", maxArea(height2)) // expected output: 1
}

func maxArea(height []int) int {
	maxArea := 0
	left, right := 0, len(height)-1

	for left < right {
		currentArea := min_Area(height[left], height[right]) * (right - left)
		maxArea = max_Area(maxArea, currentArea)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func min_Area(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max_Area(a, b int) int {
	if a > b {
		return a
	}
	return b
}