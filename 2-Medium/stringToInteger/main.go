package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	s := "-91283472332"
	fmt.Print(myAtoi(s))
}

/*
The algorithm for myAtoi(string s) is as follows:

1. Whitespace: Ignore any leading whitespace (" ").
2. Signedness: Determine the sign by checking if the next character is '-' or '+', assuming positivity if neither present.
3. Conversion: Read the integer by skipping leading zeros until a non-digit character is encountered or the end of the string is reached.
   If no digits were read, then the result is 0.
4. Rounding: If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then round the integer to remain in the range.
   Specifically, integers less than -231 should be rounded to -231, and integers greater than 231 - 1 should be rounded to 231 - 1.

Return the integer as the final result.
*/

func myAtoi(s string) int {
	isNegative := false
	var strBuilder strings.Builder
	started := false

	for _, v := range s {
		if !started {
			if unicode.IsSpace(v) {
				continue
			} else if v == '-' {
				isNegative = true
				started = true
				continue
			} else if v == '+' {
				started = true
				continue
			} else if unicode.IsNumber(v) {
				started = true
				strBuilder.WriteRune(v)
			} else {
				break
			}
		} else {
			if unicode.IsNumber(v) {
				strBuilder.WriteRune(v)
			} else {
				break
			}
		}
	}

	if strBuilder.Len() == 0 {
		return 0
	}

	numStr := strBuilder.String()
	num64, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		// If overflow, clamp to limits
		if isNegative {
			return -2147483648
		}
		return 2147483647
	}

	if isNegative {
		num64 = -num64
	}

	const (
		IntMax = 2147483647
		IntMin = -2147483648
	)

	if num64 < IntMin {
		return IntMin
	}
	if num64 > IntMax {
		return IntMax
	}

	return int(num64)
}
