package main

func minSubArrayLen(s int, nums []int) int {
	minNum := 0
	sum := 0
	length := len(nums)
	for start, end := 0, 0; end < length; end ++ {
		sum += nums[end]
		if sum >= s {
			if minNum == 0 {
				minNum = end - start + 1
			}
			for start < end {
				if sum - nums[start] >= s {
					sum -= nums[start]
					start ++
					if minNum > end - start + 1 {
						minNum = end - start + 1
					}
				} else {
					break
				}
			}
		}
	}
	return minNum
}
