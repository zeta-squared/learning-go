package main

func rotateArray(nums []int, k int) {
	k = k % len(nums)
	copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
}
