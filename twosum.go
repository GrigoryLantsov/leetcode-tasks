package main
import "fmt"
func twosum_print() {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if index, ok := numMap[complement]; ok {
			return []int{index, i}
		}
		numMap[num] = i
	}
	return []int{-1, -1}
}