package main

func maximumProduct(nums []int) int {
	var index int
	length := len(nums)
	for i := 0; i < 3; i ++ {
		index = 0
		for j := 1; j < length - i; j ++  {
			if nums[index] < nums[j] {
				index = j
			}
		}
		nums[length - i - 1], nums[index] = nums[index], nums[length - i - 1]
	}
	for i := 0; i < 2; i ++ {
		index = i
		for j := 1 + i; j < length - 3; j ++  {
			if nums[index] > nums[j] {
				index = j
			}
		}
		nums[i], nums[index] = nums[index], nums[i]
	}
	max1 := nums[length - 1] * nums[length - 2] * nums[length - 3]
	max2 := nums[0] * nums[1] * nums[length - 1]
	if max1 > max2 {
		return max1
	} else {
		return max2
	}
}