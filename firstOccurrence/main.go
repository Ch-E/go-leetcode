package main

import "fmt"

func main() {
	haystack := "abc"
	needle := "c"

	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	if (len(haystack) == 1 && len(needle) == 1) && haystack == needle {
		return 0
	}

	for i := 0; i < len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}
