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
func compress(chars []byte) int {
	readIndex, writeIndex := 0, 0
	for readIndex < len(chars) {
		currentChar := chars[readIndex]
		currentCount := 0

		// Count the number of consecutive repeating characters
		for readIndex < len(chars) && currentChar == chars[readIndex] {
			readIndex++
			currentCount++
		}

		// Write the compression result to chars
		chars[writeIndex] = currentChar
		writeIndex++

		if currentCount > 1 {
			count := strconv.Itoa(currentCount)
			for _, i := range count {
				chars[writeIndex] = byte(i)
				writeIndex++
			}
		}
	}

	return writeIndex
}
