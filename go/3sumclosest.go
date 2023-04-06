package main
import (
	"sort"
	"fmt"
)
func threeSumClosestInit() {
	sum1 := []int{-1,2,1,-4}
	fmt.Println("Input:", sum1)
	fmt.Println("Output:", threeSumClosest(sum1, 1))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	closest := nums[0] + nums[1] + nums[2]

	for i := 0; i < n-2; i++ {
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if abs(sum-target) < abs(closest-target) {
				closest = sum
			}
			if sum == target {
				return sum
			} else if sum < target {
				j++
			} else {
				k--
			}
		}
	}

	return closest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}