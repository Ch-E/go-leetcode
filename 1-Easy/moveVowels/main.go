package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "IceCreAm"
	// Output: "AceCreIm"
	fmt.Println(reverseVowels(s))
}

/*
Given a string s, reverse only all the vowels in the string and return it.
The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.
*/
func reverseVowels(s string) string {
	runes := []rune(s)

	i := 0
	j := len(runes) - 1

	for i < j {
		for i < j && !strings.ContainsRune("aeiouAEIOU", runes[i]) {
			i++
		}
		for i < j && !strings.ContainsRune("aeiouAEIOU", runes[j]) {
			j--
		}

		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}

	return string(runes)
}
