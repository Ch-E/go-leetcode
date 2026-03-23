package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
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

	small := nums[0]
	medium := nums[1]
	large := nums[2]

	for i := 1; i < len(nums); i++ {
		if small >= nums[i] {
			small = nums[i]
		} else if medium >= nums[i] {
			medium = nums[i]
		} else {
			large = nums[i]
		}
	}

	if small < medium && medium < large {
		return true
	}

	/*
		for i := 0; i < len(nums); i++ {
			for j := 0; j < len(nums); j++ {
				for k := 0; k < len(nums); k++ {
					if i < j && j < k && nums[i] < nums[j] && nums[j] < nums[k] {
						return true
					}
				}
			}
		}
	*/

	return false
}
