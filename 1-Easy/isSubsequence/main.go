package main

import "fmt"

func main() {
	s := "aza"
	t := "abzba"

	fmt.Println(isSubsequence(s, t))
}

/*
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
A subsequence of a string is a new string that is formed from the original string by
deleting some (can be none) of the characters without disturbing the relative positions
of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).
*/
func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	j := 0
	matchingCount := 0

	for i := 0; i < len(s); i++ {
		for j < len(t) {
			if s[i] != t[j] {
				j++
			} else {
				j++
				matchingCount++
				break
			}
		}
	}

	return matchingCount == len(s)
}
