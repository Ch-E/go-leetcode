package main

import "fmt"

func main() {
	nums := []int{2, 3, 4, 4, 4, 5, 5, 6}
	fmt.Print(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	idx := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[idx] {
			idx++
			nums[idx] = nums[i]
		}
	}

	return idx + 1
}
