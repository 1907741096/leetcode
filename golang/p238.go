package main

func productExceptSelf(nums []int) []int {
	leftArr := make([]int, len(nums))
	rightArr := make([]int, len(nums))
	for i := 0; i < len(nums); i ++ {
		if i == 0 {
			leftArr[i] = nums[i]
		} else {
			leftArr[i] = leftArr[i - 1] * nums[i]
		}
	}
	for i := len(nums) - 1; i >= 0; i -- {
		if i == len(nums) - 1 {
			rightArr[i] = nums[i]
		} else {
			rightArr[i] = rightArr[i + 1] * nums[i]
		}
	}
	arr := make([]int, len(nums))
	arr[0] = rightArr[1]
	for i := 1; i < len(nums) - 1; i ++ {
		arr[i] = leftArr[i - 1] * rightArr[i + 1]
	}
	arr[len(nums) - 1] = leftArr[len(nums) - 2]
	return arr
}