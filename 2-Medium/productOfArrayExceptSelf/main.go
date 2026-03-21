package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}

/*
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
You must write an algorithm that runs in O(n) time and without using the division operation.
*/
func productExceptSelf(nums []int) []int {
	result := make([]int, 0)
	product := 1

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j != i {
				product *= nums[j]
			}
		}

		result = append(result, product)
		product = 1
	}

	return result
}
