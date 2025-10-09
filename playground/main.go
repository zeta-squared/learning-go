package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	x := []int{1, 2, 3, 0, 0, 0}
	merge(x, []int{2, 5, 6}, 3, 3)
	fmt.Println(x)
}

func merge(nums1, nums2 []int, m, n int) {
	temp := nums1[:m]
	i, j := 0, 0
	for k := range len(nums1) {
		for temp[i] <= nums2[j] && i < m && j < n {
			nums1[k] = temp[i]
			i++
		}

		if nums1[i] < nums2[j] && i < m && j < n {
			nums1[k] = nums2[j]
			j++
		}
	}

	for v, i := range temp {
		nums1[i] = v
	}
}
