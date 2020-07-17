package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	flag := false
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			flag = true
			left = mid
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	count := 0
	if flag {
		l, r := left, left + 1
		for l >= 0 && nums[l] == target {
			count ++
			l --
		}
		for r < len(nums) && nums[r] == target {
			count ++
			r ++
		}
	}
	return count
}