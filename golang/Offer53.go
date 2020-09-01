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

func missingNumber(nums []int) int {
	for i := 1; i < len(nums); i += 2 {
		if nums[i] - 2 == nums[i - 1] {
			return nums[i] - 1
		}
		if i + 1 < len(nums) && nums[i] + 2 == nums[i + 1] {
			return nums[i] + 1
		}
	}
	if nums[0] != 0 {
		return 0
	} else if len(nums) == 1{
		return nums[0] + 1
	} else if nums[0] + 2 == nums[1] {
		return nums[0] + 1
	} else {
		return nums[len(nums) - 1] + 1
	}
}