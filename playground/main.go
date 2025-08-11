package main

import (
	"fmt"
)

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
	s = "bbbbb"
	fmt.Println(lengthOfLongestSubstring(s))
	s = "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
	s = " "
	fmt.Println(lengthOfLongestSubstring(s))
	s = "au"
	fmt.Println(lengthOfLongestSubstring(s))
	s = "au"
	fmt.Println(lengthOfLongestSubstring(s))
}
