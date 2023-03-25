package main
import "fmt"

func findMedianInit() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	median := findMedianSortedArrays(nums1, nums2)
	fmt.Printf("Median: %f", median)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		nums1, nums2, m, n = nums2, nums1, n, m
	}
	imin, imax, half := 0, m, (m+n+1)/2
	for imin <= imax {
		i := (imin + imax) / 2
		j := half - i
		if i < m && nums2[j-1] > nums1[i] {
			imin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			imax = i - 1
		} else {
			var maxLeft, minRight int
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = findMax(nums1[i-1], nums2[j-1])
			}
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = findMin(nums1[i], nums2[j])
			}
			return float64(maxLeft+minRight) / 2.0
		}
	}
	return 0.0
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}