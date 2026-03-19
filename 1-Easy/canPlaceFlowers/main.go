package main

import "fmt"

func main() {
	flowerbed := []int{0, 0, 1, 0, 1}
	n := 1

	fmt.Println(canPlaceFlowers(flowerbed, n))
}

/*
You have a long flowerbed in which some of the plots are planted, and some are not.
However, flowers cannot be planted in adjacent plots.
Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty,
and an integer n, return true if n new flowers can be planted in the flowerbed
without violating the no-adjacent-flowers rule and false otherwise.
*/
func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 1 && len(flowerbed) == 1 && flowerbed[0] == 0 {
		return true
	}
	if n == 1 && len(flowerbed) == 2 && flowerbed[0] == 0 && flowerbed[1] == 0 {
		return true
	}
	if n > 0 {
		for i := 1; i < len(flowerbed)-1; i++ {
			if len(flowerbed) > 1 && flowerbed[0] == 0 && flowerbed[1] == 0 {
				flowerbed[0] = 1
				n--
			}
			if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
			}
			if flowerbed[len(flowerbed)-1] == 0 && flowerbed[len(flowerbed)-2] == 0 {
				flowerbed[len(flowerbed)-1] = 1
				n--
			}
		}
	}

	if n > 0 {
		return false
	}

	return true
}
