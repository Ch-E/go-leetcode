package main

import "fmt"

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3

	fmt.Println(kidsWithCandies(candies, extraCandies))
}

/*
There are n kids with candies. You are given an integer array candies,
where each candies[i] represents the number of candies the ith kid has,
and an integer extraCandies, denoting the number of extra candies that you have.

Return a boolean array result of length n, where result[i] is true if,
after giving the ith kid all the extraCandies, they will have the greatest
number of candies among all the kids, or false otherwise.

Note that multiple kids can have the greatest number of candies.
*/

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, 0)

	for i := range candies {
		isGreatest := true
		sumCandies := candies[i] + extraCandies

		for j := range candies {
			if sumCandies < candies[j] {
				isGreatest = false
				break
			}
		}

		result = append(result, isGreatest)
	}

	return result
}

/*
Optimized solution O(n):
func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0

	for i := 0; i < len(candies); i++ {
		if max < candies[i] {
			max = candies[i]
		}
	}

	result := make([]bool, len(candies))

	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= max {
			result[i] = true
		}
	}

	return result
}
*/
