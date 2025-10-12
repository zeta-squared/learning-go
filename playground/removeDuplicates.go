package main

func removeDuplicates(nums []int) int {
	unique := nums[0]
	rIndex := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != unique {
			nums[rIndex] = nums[i]
			unique = nums[i]
			rIndex++
		}
	}

	return rIndex
}
