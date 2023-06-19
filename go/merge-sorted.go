package main
import "fmt"

func mergeSortedInit() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	expected := []int{1, 2, 2, 3, 5, 6}

	merge(nums1, m, nums2, n)

	// Check that nums1 now contains the expected values
	for i, v := range nums1 {
		if v != expected[i] {
			fmt.Printf("Error: expected %v but got %v\n", expected, nums1)
			return
		}
	}

	fmt.Println("Test passed")
}


//We iterate as long as both arrays have elements remaining. If nums1[i] is larger than nums2[j], we place it at nums1[k], decrement i, and decrement k. Otherwise, we place nums2[j] at nums1[k], decrement j, and decrement k.
//After the first loop, if there are any elements remaining in nums2, we copy them to nums1 starting at k and decrementing k and j until we've copied all the elements.
//This algorithm runs in O(m + n) time because we only iterate over each element in both arrays once.
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
