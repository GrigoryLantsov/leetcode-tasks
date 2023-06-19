package main

import (
	"fmt"
)

func largestAltitude(gain []int) int {
	maxAltitude, altitude := 0, 0
	for _, g := range gain {
		altitude += g
		if altitude > maxAltitude {
			maxAltitude = altitude
		}
	}
	return maxAltitude
}

func largestAltitudeInit() {
	gain1 := []int{-5, 1, 5, 0, -7}
	fmt.Println(largestAltitude(gain1))  // Output: 1

	gain2 := []int{-4, -3, -2, -1, 4, 3, 2}
	fmt.Println(largestAltitude(gain2))  // Output: 0
}