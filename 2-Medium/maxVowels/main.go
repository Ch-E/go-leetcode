package main

import "fmt"

func main() {
	s := "abciiidef"
	k := 3

	fmt.Println(maxVowels(s, k))
}

/*
Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.
Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.
*/
func maxVowels(s string, k int) int {
	left, right := 0, 0
	result := 0
	vowel := "aeiou"

	for right < len(s) {
		vowelCount := 0
		window := right - left + 1

		if window == k {
			// Check for vowels

			left++
		}

		if result < vowelCount {
			result = vowelCount
		}

		right++
	}

	return result
}
