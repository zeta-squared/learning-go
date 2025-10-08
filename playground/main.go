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
	temp := make([]int, m+n)
	i, j := 0, 0
	for k := range len(nums1) {
		if i > m {
			temp[k] = nums2[j]
			j++
		} else if j > n {
			temp[k] = nums1[i]
			i++
		}

		temp[k] = min(nums1[i], nums2[j])
		if nums1[i] <= nums2[j] {
			i++
		} else {
			j++
		}
	}
}
