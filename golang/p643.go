package main

func findMaxAverage(nums []int, k int) float64 {
	maxNum, nowNum := 0, 0
	for i := 0; i < k; i ++ {
		nowNum += nums[i]
	}
	maxNum = nowNum
	for i := k; i < len(nums); i ++ {
		nowNum += nums[i]
		nowNum -= nums[i - 1]
		if maxNum < nowNum {
			maxNum = nowNum
		}
	}
	return float64(maxNum) / float64(k)
}