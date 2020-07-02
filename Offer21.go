package main

func exchange(nums []int) []int {
	left, right := 0, len(nums)
	for left < right {
		if nums[left] % 2 == 0 && nums[right] % 2 == 1 {
			temp := nums[left]
			nums[left] = nums[right]
			nums[right] = temp
			left ++
			right --
		}
		if nums[left] % 2 == 1 {
			left ++
		}
		if nums[right] % 2 == 0 {
			right --
		}
	}
	return nums
}