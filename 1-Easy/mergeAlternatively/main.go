package main

import "fmt"

func main() {
	word1 := "abc"
	word2 := "pqrst"

	fmt.Println(mergeAlternately(word1, word2))
}

/*
You are given two strings word1 and word2.
Merge the strings by adding letters in alternating order, starting with word1.
If a string is longer than the other, append the additional letters onto the end of the merged string.
*/
func mergeAlternately(word1 string, word2 string) string {
	result := []byte{}

	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			result = append(result, word1[i])
		}

		if i < len(word2) {
			result = append(result, word2[i])
		}
	}

	return string(result)
}
