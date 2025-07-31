package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
	nums = []int{0, 0, 0, 0}
	fmt.Println(threeSum(nums))
	nums = []int{-2, 0, 1, 1, 2}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	var trips [][]int
	sort.Ints(nums)
	fmt.Println(nums)
	for i, v1 := range nums {
		if i >= len(nums)-2 { // Not possible to form any more triplets.
			break
		}

		k := len(nums) - 1
		for j := i + 1; j < k; {
			v2 := nums[j]
			v3 := nums[k]

			if v1+v2+v3 < 0 {
				j++
			} else if v1+v2+v3 > 0 {
				k--
			} else {
				if len(trips) > 0 && !slices.Equal(trips[len(trips)-1], []int{v1, v2, v3}) {
					trips = append(trips, []int{v1, v2, v3})
				}
			}
		}
	}

	return trips
}
