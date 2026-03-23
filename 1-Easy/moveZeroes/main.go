package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

/*
Given an integer array nums, move all 0's to the end of it while maintaining
the relative order of the non-zero elements.
Note that you must do this in-place without making a copy of the array.
*/
func moveZeroes(nums []int) {
	i := 0
	j := len(nums) - 1

	for i < j {
		if nums[i] == 0 && nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		} else {
			i++
		}
	}
}

/*
func moveZeroes(nums []int)  {
    left := 0
    for right := 0; right < len(nums); right++ {
        if nums[right] != 0 {
            nums[left], nums[right] = nums[right], nums[left]
            left++
        }
    }
}
*/
