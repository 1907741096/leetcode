package main

func subarraySum(nums []int, k int) int {
	countSum := 0
	countMap := make(map[int]int)
	countMap[0] = 1
	arrSum := make([]int, len(nums))
	for i := 0; i < len(nums); i ++ {
		if i == 0 {
			arrSum[i] = nums[i]
		} else {
			arrSum[i] = arrSum[i - 1] + nums[i]
		}
		if count, ok := countMap[arrSum[i] - k]; ok {
			countSum += count
		}
		if _, ok := countMap[arrSum[i]]; ok {
			countMap[arrSum[i]] ++
		} else {
			countMap[arrSum[i]] = 1
		}
	}
	return countSum
}