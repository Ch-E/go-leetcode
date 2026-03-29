package main

import "fmt"

func main() {
	nums := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	fmt.Println(longestSubarray(nums))
}

/*
Given a binary array nums, you should delete one element from it.
Return the size of the longest non-empty subarray containing only 1's in the resulting array.
Return 0 if there is no such subarray.
*/
func longestSubarray(nums []int) int {
	left, right := 0, 0
	k := 1

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

	return right - left - 1
}
