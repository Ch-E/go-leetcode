package main

import (
	"fmt"
	"slices"
)

func main() {
	s := "bbbbbbbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	storage := make(map[int][]byte)

	for i := 0; i < len(s)-1; i++ {
		// {0, a}
		storage[i] = []byte{s[i]}
		// Compare a with b
		// TODO: Compare b with the previous letters in the string
		// {1, b}
		for j := i + 1; j < len(s); j++ {
			if slices.Contains(storage[i], s[j]) {
				break
			} else {
				storage[i] = append(storage[i], s[j])
			}
		}
	}

	maxLength := 0
	for _, v := range storage {
		if len(v) > maxLength {
			maxLength = len(v)
		}
	}

	return maxLength
}
