package main

import (
	"fmt"
	"slices"
)

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
	result := make([]float64, 0)

	for i := 0; i <= len(nums)-k; i++ {
		j := i + 1
		sum := nums[i]
		for j < i+k {
			sum += nums[j]
			j++
		}

		avg := float64(sum) / float64(k)
		result = append(result, float64(avg))
	}

	return slices.Max(result)
}
