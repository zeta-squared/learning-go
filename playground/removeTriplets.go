package main

func removeTriplets(nums []int) int {
	unique := nums[0]
	count := 1
	rIndex := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != unique {
			nums[rIndex] = nums[i]
			unique = nums[i]
			count = 1
			rIndex++
		} else {
			count++
			if count <= 2 {
				nums[rIndex] = nums[i]
				rIndex++
				count++
			}
		}
	}

	return rIndex
}
