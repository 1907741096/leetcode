package main

import (
	"sort"
)

func isStraight(nums []int) bool {
	sort.Ints(nums)
	min, max := 0, 0
	for i := 0; i < len(nums); i ++ {
		if nums[i] != 0 {
			if i >= 1 && nums[i] == nums[i - 1] {
				return false
			} else {
				if min == 0 || min > nums[i] {
					min = nums[i]
				}
				if max == 0 || max < nums[i] {
					max = nums[i]
				}
			}
		}
	}
	return max - min <= 5
}