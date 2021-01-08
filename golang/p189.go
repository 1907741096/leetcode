package main

func rotate(nums []int, k int)  {
	reserve(nums)
	reserve(nums[:k % len(nums)])
	reserve(nums[k % len(nums):])
}

func reserve(nums []int) {
	var temp int
	for i := 0; i < len(nums) / 2; i ++ {
		temp = nums[i]
		nums[i] = nums[len(nums) - i - 1]
		nums[len(nums) - i - 1] = temp
	}
}
