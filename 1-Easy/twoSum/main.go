package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 5, 2, 6}
	fmt.Println(twoSum(nums, 7))
}

func twoSum(nums []int, target int) []int {
	mapStore := make(map[int]int)

	for i, v := range nums {
		valNeeded := target - v
		// Check if the value we need is in the map - if it exists, indexNeeded would be the index of that value
		if indexNeeded, ok := mapStore[valNeeded]; ok {
			return []int{indexNeeded, i}
		} else {
			// Store the value in the map
			mapStore[v] = i
		}
	}

	return nil
}
