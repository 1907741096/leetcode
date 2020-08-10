package main

func majorityElement(nums []int) int {
	numSet := make(map[int]int)
	max := 0
	maxNum := 0
	for _, num := range nums {
		if _, ok := numSet[num]; ok {
			numSet[num] ++
		} else {
			numSet[num] = 1
		}
		if max < numSet[num] {
			max = numSet[num]
			maxNum = num
		}
	}
	return maxNum
}
