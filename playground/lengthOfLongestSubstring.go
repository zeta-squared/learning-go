package main

func lengthOfLongestSubstring(s string) int {
	max := 0
	for i := range len(s) {
		strMap := map[byte]bool{
			s[i]: true,
		}
		curr := 1
		for j := i + 1; j < len(s); j++ {
			if _, ok := strMap[s[j]]; ok {
				break
			}

			strMap[s[j]] = true
			curr++
		}

		if curr > max {
			max = curr
		}
	}

	return max
}
