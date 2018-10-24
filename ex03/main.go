package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	location := [256]int{}
	for i := range location {
		location[i] = -1
	}

	maxlen, left := 0, 0

	for i := 0; i < len(s); i++ {
		if location[s[i]] >= left {
			left = location[s[i]] + 1
		} else if i+1-left > maxlen {
			maxlen = i + 1 - left
		}
		location[s[i]] = i
	}

	return maxlen
}

func main() {
	s := "aabbabc"
	fmt.Println("longest sub string is ", lengthOfLongestSubstring(s))
}
