package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2

	fmt.Println(longestOnes(nums, k))
}

/*
Given a binary array nums and an integer k, return the maximum number of
consecutive 1's in the array if you can flip at most k 0's.
*/
func longestOnes(nums []int, k int) int {
	left, right := 0, 0

	for right < len(nums) {
		if nums[right] == 0 {
			k--
		}
		right++

		if k < 0 {
			if nums[left] == 0 {
				k++
			}
			left++
		}
	}

	return right - left
}
