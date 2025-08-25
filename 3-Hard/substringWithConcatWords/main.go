// https://leetcode.com/problems/substring-with-concatenation-of-all-words/description/
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	fmt.Println(findSubstring(s, words))
}

func findSubstring(s string, words []string) []int {
	perms := generatePermutations(words)
	var wordBuilder strings.Builder
	concatWord := []string{}

	for i := 0; i < len(perms); i++ {
		wordBuilder.Reset()
		for j := 0; j < len(perms[i]); j++ {
			wordBuilder.WriteString(perms[i][j])
		}
		concatWord = append(concatWord, wordBuilder.String())
	}

	wordIndex := []int{}

	for i := 0; i < len(concatWord); i++ {
		word := concatWord[i]
		for j := 0; j <= len(s)-len(word); j++ {
			if s[j:j+len(word)] == word {
				wordIndex = append(wordIndex, j)
			}
		}
	}

	// Remove duplicates
	unique := make(map[int]struct{})
	for _, idx := range wordIndex {
		unique[idx] = struct{}{}
	}
	result := make([]int, 0, len(unique))
	for idx := range unique {
		result = append(result, idx)
	}

	return result
}

func generatePermutations(words []string) [][]string {
	if len(words) == 0 {
		return [][]string{}
	}

	// Base case
	if len(words) == 1 {
		return [][]string{words}
	}

	var result [][]string

	for i, word := range words {
		// Create new slice without the current word
		remaining := append([]string{}, words[:i]...)
		remaining = append(remaining, words[i+1:]...)

		// Generate all permutations of the remaining words
		for _, perm := range generatePermutations(remaining) {
			result = append(result, append([]string{word}, perm...))
		}
	}

	return result
}
