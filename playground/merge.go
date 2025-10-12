package main

func merge(nums1, nums2 []int, m, n int) {
	temp := make([]int, m)
	copy(temp, nums1[:m])
	i, j := 0, 0
	for i < m && j < n {
		if temp[i] <= nums2[j] {
			nums1[i+j] = temp[i]
			i++
		} else {
			nums1[i+j] = nums2[j]
			j++
		}
	}

	for ; i < m; i++ {
		nums1[i+j] = temp[i]
	}

	for ; j < n; j++ {
		nums1[i+j] = nums2[j]
	}
}
