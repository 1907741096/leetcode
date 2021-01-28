package main

func pivotIndex(nums []int) int {
	res := -1
	if len(nums) == 0 {
		return res
	}
	leftArr := make([]int, len(nums) + 1)
	rightArr := make([]int, len(nums) + 1)

	for i := 0; i < len(nums); i ++ {
		leftArr[i + 1] = leftArr[i] + nums[i]
		rightArr[len(nums) - i - 1] = rightArr[len(nums) - i] + nums[len(nums) - i - 1]
	}
	for i := 0; i < len(nums); i ++ {
		if leftArr[i] == rightArr[i + 1] {
			res = i
			break
		}
	}
	return res
}