package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	trips := [][]int{}
	tripsMap := map[string]bool{}
	sort.Ints(nums)
	for i := range nums {
		if i >= len(nums)-2 { // Not possible to form any more triplets.
			break
		}

		k := len(nums) - 1
		j := i + 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			switch {
			case sum == 0:
				trip := []int{nums[i], nums[j], nums[k]}
				tripsString := fmt.Sprint(trip)
				if _, ok := tripsMap[tripsString]; !ok {
					tripsMap[tripsString] = true
					trips = append(trips, trip)
				}

				j++
				for j < k && nums[j-1] == nums[j] {
					j++
				}

				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			case sum > 0:
				k--
			default:
				j++
			}
		}
	}

	return trips
}
