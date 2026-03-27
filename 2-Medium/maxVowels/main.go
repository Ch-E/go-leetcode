package main

import (
	"fmt"
	"strings"
)

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
	vowelCount := 0

	for right < len(s) {
		window := right - left + 1

		if strings.IndexByte(vowel, s[right]) != -1 {
			vowelCount++
		}

		if window == k {
			if result < vowelCount {
				result = vowelCount
			}

			if strings.IndexByte(vowel, s[left]) != -1 {
				vowelCount--
			}

			left++
		}

		right++
	}

	return result
}
