package main

import (
	"fmt"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(sortStrings(strs))
	fmt.Println(groupAnagrams(strs))
	strs = []string{"a"}
	fmt.Println(groupAnagrams(strs))
	strs = []string{""}
	fmt.Println(groupAnagrams(strs))
	strs = []string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"}
	fmt.Println(groupAnagrams(strs))
}
