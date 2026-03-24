package main

import (
	"fmt"
)

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}

/*
You are given an integer array height of length n. There are n vertical lines drawn such that
the two endpoints of the ith line are (i, 0) and (i, height[i]).
Find two lines that together with the x-axis form a container, such that the container contains the most water.
Return the maximum amount of water a container can store.
*/
func maxArea(height []int) int {
	area := 0
	maxArea := 0

	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height) && j != i; j++ {
			area = (j - i) * min(height[i], height[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}
