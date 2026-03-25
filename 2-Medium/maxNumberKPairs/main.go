package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	k := 5

	fmt.Println(maxOperations(nums, k))
}

/*
You are given an integer array nums and an integer k.
In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.
Return the maximum number of operations you can perform on the array.
*/
func maxOperations(nums []int, k int) int {
	result := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if k-nums[i] == nums[j] {
				result++
				nums = append(nums[:j], nums[j+1:]...)
				nums = append(nums[:i], nums[i+1:]...)
				i--
				break
			}
		}
	}

	return result
}
