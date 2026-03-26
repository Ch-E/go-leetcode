package main

import "fmt"

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4

	fmt.Println(findMaxAverage(nums, k))
}

/*
You are given an integer array nums consisting of n elements, and an integer k.
Find a contiguous subarray whose length is equal to k that has the maximum average value
and return this value. Any answer with a calculation error less than 10^-5 will be accepted.
*/
func findMaxAverage(nums []int, k int) float64 {
	j := 1

	for i := 0; i < len(nums); i++ {
		for j < k && j < len(nums) {

		}
	}
}
