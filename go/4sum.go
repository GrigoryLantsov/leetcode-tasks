//we can use a similar approach to the three sum problem. We can sort the array, then fix the first two elements and use two pointers to search for the other two elements that sum up to the target. We can use a set to keep track of unique quadruplets.

package main
import (
	"sort"
	"fmt"
)
func fourSumInit() {
	sum1 := []int{1,0,-1,0,-2,2}
	fmt.Println("Input:", sum1)
	fmt.Println("Output:", fourSum(sum1, 0))
}

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	n := len(nums)
	if n < 4 {
		return result
	}

	sort.Ints(nums)

	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i+1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left := j+1
			right := n-1

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]

				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--

					for left < right && nums[left] == nums[left-1] {
						left++
					}

					for left < right && nums[right] == nums[right+1] {
						right--
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return result
}

