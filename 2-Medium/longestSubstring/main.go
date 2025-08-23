package main

import (
	"fmt"
	"slices"
)

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	storage := make(map[int][]byte)

	for i := 0; i < len(s)-1; i++ {
		storage[i] = []byte{s[i]}

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

// * Optimized solution (sliding window):
/*
func lengthOfLongestSubstring(s string) int {
    n := len(s)
    maxLength := 0
    lastIndex := make([]int, 128)

    start := 0
    for end := 0; end < n; end++ {
        currentChar := s[end]
        if lastIndex[currentChar] > start {
            start = lastIndex[currentChar]
        }
        if end-start+1 > maxLength {
            maxLength = end - start + 1
        }
        lastIndex[currentChar] = end + 1
    }

    return maxLength
}
*/
