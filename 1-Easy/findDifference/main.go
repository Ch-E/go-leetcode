package main

import (
	"fmt"
	"slices"
)

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}

	fmt.Println(findDifference(nums1, nums2))
}

/*
Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:
answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
answer[1] is a list of all distinct integers in nums2 which are not present in nums1.
Note that the integers in the lists may be returned in any order.
*/
func findDifference(nums1 []int, nums2 []int) [][]int {
	uniqueNums1 := make([]int, 0)
	uniqueNums2 := make([]int, 0)

	for i := 0; i < len(nums1); i++ {
		isMatch := false
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				isMatch = true
			}
		}

		if !isMatch {
			uniqueNums1 = append(uniqueNums1, nums1[i])
		}
	}

	for i := 0; i < len(nums2); i++ {
		isMatch := false
		for j := 0; j < len(nums1); j++ {
			if nums2[i] == nums1[j] {
				isMatch = true
			}
		}

		if !isMatch {
			uniqueNums2 = append(uniqueNums2, nums2[i])
		}
	}

	uniqueNums1 = slices.Compact(uniqueNums1)
	uniqueNums2 = slices.Compact(uniqueNums2)

	return [][]int{uniqueNums1, uniqueNums2}
}
