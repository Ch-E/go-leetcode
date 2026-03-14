package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0, 2}

	sortColors(nums)
	fmt.Println(nums)
}

// Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects
// of the same color are adjacent, with the colors in the order red, white, and blue..
func sortColors(nums []int) {
	colorCount := []int{0, 0, 0}

	for _, v := range nums {
		switch v {
		case 0:
			colorCount[0]++
		case 1:
			colorCount[1]++
		default:
			colorCount[2]++
		}
	}

	for i := range nums {
		if colorCount[0] > 0 {
			nums[i] = 0
			colorCount[0]--
		} else if colorCount[1] > 0 {
			nums[i] = 1
			colorCount[1]--
		} else {
			nums[i] = 2
			colorCount[2]--
		}
	}
}
