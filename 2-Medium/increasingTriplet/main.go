package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{0, 4, 2, 1, 0, -1, -3}
	fmt.Println(increasingTriplet(nums))
}

/*
Given an integer array nums, return true if there exists a triple of indices (i, j, k)
such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.
*/
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	small := math.MaxInt64
	medium := math.MaxInt64

	for i := range nums {
		if nums[i] <= small {
			small = nums[i]
		} else if nums[i] <= medium {
			medium = nums[i]
		} else {
			return true
		}
	}

	return false
}
