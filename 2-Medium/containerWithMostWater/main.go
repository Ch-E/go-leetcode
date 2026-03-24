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

	i, j := 0, len(height)-1

	for i < j {
		area = (j - i) * min(height[i], height[j])

		if area > maxArea {
			maxArea = area
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return maxArea
}
