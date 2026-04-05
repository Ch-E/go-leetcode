package main

import (
	"fmt"
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
	seen1 := make(map[int]bool)
	seen2 := make(map[int]bool)
	unique1 := make([]int, 0)
	unique2 := make([]int, 0)

	for _, v := range nums1 {
		seen1[v] = true
	}

	for _, v := range nums2 {
		seen2[v] = true
	}

	for i := range seen1 {
		if !seen2[i] {
			unique1 = append(unique1, i)
		}
	}

	for i := range seen2 {
		if !seen1[i] {
			unique2 = append(unique2, i)
		}
	}

	return [][]int{unique1, unique2}
}
