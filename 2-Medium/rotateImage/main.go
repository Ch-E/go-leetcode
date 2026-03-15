package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)

	fmt.Print(matrix)
}

/*
You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
You have to rotate the image in-place, which means you have to modify the input 2D matrix directly.
DO NOT allocate another 2D matrix and do the rotation.

Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: 		[[7,4,1],[8,5,2],[9,6,3]]
*/
func rotate(matrix [][]int) {
	n := len(matrix)

	for r := 0; r < n; r++ {
		for c := r + 1; c < n; c++ {
			matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
		}
	}

	for r := 0; r < n; r++ {
		for c := 0; c < n/2; c++ {
			matrix[r][c], matrix[r][n-1-c] = matrix[r][n-1-c], matrix[r][c]
		}
	}
}
