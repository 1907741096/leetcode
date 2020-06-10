package main

func rob(nums []int) int {
	max := 0
	for i := 0; i <len(nums); i ++ {
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
