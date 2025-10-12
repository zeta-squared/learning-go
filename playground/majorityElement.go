package main

func majorityElement(nums []int) int {
	values := make(map[int]int)
	for _, v := range nums {
		values[v]++
		if values[v] > len(nums)/2 {
			return v
		}
	}

	return 0
}
