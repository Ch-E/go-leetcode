/*
List of Analogous Backtracking Questions:
https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
https://leetcode.com/problems/n-queens/description/
https://leetcode.com/problems/permutations-ii/description/
https://leetcode.com/problems/combination-sum-ii/description/
https://leetcode.com/problems/permutations/description/
https://leetcode.com/problems/generate-parentheses/description/
https://leetcode.com/problems/subsets/description/
*/

// https://leetcode.com/problems/combination-sum/solutions/5976264/complex-backtracking-interview-prepare-list-of-backtracking-questions-beats-100/
package main

import "fmt"

func main() {
	candidates := []int{2, 3, 5}
	target := 8

	fmt.Println(combinationSum(candidates, target))
}

// findCombinations is a recursive helper function to find all combinations
// index: current position in candidates array
// target: remaining sum to achieve
// candidates: input array of numbers
// current: current combination being built
// result: stores all valid combinations
func findCombinations(index, target int, candidates []int, current []int, result *[][]int) {
	// Base case: if target becomes 0, we found a valid combination
	if target == 0 {
		// Create a copy of current combination and add to result
		combination := make([]int, len(current))
		copy(combination, current)
		*result = append(*result, combination)
		return
	}

	// Try all possible numbers from current index
	for i := index; i < len(candidates); i++ {
		// Only proceed if current number doesn't exceed target
		if candidates[i] <= target {
			// Include current number in combination
			current = append(current, candidates[i])

			// Recursive call with:
			// - same index i (allowing reuse of same number)
			// - reduced target by current number
			findCombinations(i, target-candidates[i], candidates, current, result)

			// Backtrack: remove the last added number to try other combinations
			current = current[:len(current)-1]
		}
	}
}

// combinationSum finds all unique combinations of candidates where the numbers sum to target
func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0) // Stores all valid combinations
	current := make([]int, 0)  // Temporary slice to build combinations
	findCombinations(0, target, candidates, current, &result)
	return result
}
