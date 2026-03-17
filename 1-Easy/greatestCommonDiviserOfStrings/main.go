package main

import "fmt"

func main() {
	str1 := "ABCABC"
	str2 := "ABC"

	fmt.Println(gcdOfStrings(str1, str2))
}

/*
For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t
(i.e., t is concatenated with itself one or more times).
Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.
*/
func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	a, b := len(str1), len(str2)
	for b > 0 {
		a, b = b, a%b
	}

	return str1[:a]
}
