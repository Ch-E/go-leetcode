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
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || firstWord[i] != strs[j][i] {
				return string(prefix)
			}
		}

		prefix = append(prefix, firstWord[i])
	}

	return string(prefix)
}
