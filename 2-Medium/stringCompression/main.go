package main

import (
	"fmt"
	"strconv"
)

func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(compress(chars))
}

/*
Given an array of characters chars, compress it using the following algorithm:
Begin with an empty string s. For each group of consecutive repeating characters in chars:
If the group's length is 1, append the character to s.
Otherwise, append the character followed by the group's length.
The compressed string s should not be returned separately, but instead, be stored in the input character array chars.
Note that group lengths that are 10 or longer will be split into multiple characters in chars.
After you are done modifying the input array, return the new length of the array.
You must write an algorithm that uses only constant extra space.
Note: The characters in the array beyond the returned length do not matter and should be ignored.
*/
func compress(chars []byte) string {
	counter := 1
	compressed := ""

	for i := 0; i < len(chars); i++ {
		for j := i + 1; j < len(chars); j++ {
			if chars[i] == chars[j] {
				counter++
			} else {
				break
			}
		}

		if counter == 1 {
			compressed += string(chars[i])
		} else {
			compressed += string(chars[i]) + strconv.Itoa(counter)
		}

		i += counter - 1
		counter = 1
	}

	return compressed
}
