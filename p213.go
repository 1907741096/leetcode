package main

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else {
		return max(robMax(nums[1:]), robMax(nums[:len(nums) - 1]))
	}
}

func robMax(n []int) int {
	max := 0
	nums := make([]int, len(n))
	copy(nums, n)
	for i := 0; i < len(nums); i ++ {
		if i == 1 || i == 0 {

		} else if i == 2 {
			nums[i] += nums[i - 2]
		} else {
			if nums[i - 2]  < nums[i - 3]  {
				nums[i] += nums[i - 3]
			} else {
				nums[i] += nums[i - 2]
			}
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
