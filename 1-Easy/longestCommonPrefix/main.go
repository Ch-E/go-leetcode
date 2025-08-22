package main

import (
	"fmt"
)

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	prefix := []byte{}
	firstWord := strs[0]

	for i := 0; i < len(firstWord); i++ {
		letter := firstWord[i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != letter {
				return string(prefix)
			}
		}

		prefix = append(prefix, letter)
	}

	return string(prefix)
}
