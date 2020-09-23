package main

func maxSlidingWindow(nums []int, k int) []int {
	arr := make([]int, len(nums) - k + 1)
	max := nums[0]
	for i := 1; i < k; i ++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	arr[0] = max
	for i := k; i < len(nums); i ++ {
		if nums[i - 1] != arr[i - k] {
			max = nums[i]
			for j := i + 1; j < k; j ++ {
				if nums[i] > max {
					max = nums[i]
				}
			}
			arr[i - k + 1] = max
		} else {
			if nums[i] > arr[i - 1] {
				arr[i - k + 1] = nums[i]
			} else {
				arr[i - k + 1] = arr[i - k]
			}
		}
	}
	return arr
}