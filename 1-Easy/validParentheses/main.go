// https://leetcode.com/problems/valid-parentheses/description/
package main

import "fmt"

func main() {
	s := "([]){}"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []rune{}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, rune(s[i]))
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]   // Get top element of stack
			stack = stack[:len(stack)-1] // Remove top element of stack

			if pairs[rune(s[i])] != top {
				return false
			}
		}
	}

	return len(stack) == 0
}
