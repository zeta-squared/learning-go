package main

import "slices"

func groupAnagrams(strs []string) [][]string {
	anagrams := [][]string{}
	// Sort all strings.
	sortedStrs := sortStrings(strs)
	for i, v := range sortedStrs {
		if v == "EOF" {
			continue
		}

		subanagrams := []string{strs[i]}
		for j := i + 1; j < len(sortedStrs); j++ {
			if v == sortedStrs[j] {
				subanagrams = append(subanagrams, strs[j])
				sortedStrs[j] = "EOF"
			}
		}

		anagrams = append(anagrams, subanagrams)
	}

	return anagrams
}

func sortStrings(s []string) []string {
	sortedStrings := []string{}
	for _, v := range s {
		strBytes := []byte(v)
		slices.Sort(strBytes)
		sortedStrings = append(sortedStrings, string(strBytes))
	}

	return sortedStrings
}
