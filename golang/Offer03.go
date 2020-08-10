package main

func findRepeatNumber(nums []int) int {
	hashSet := make(map[int]int)
	for _, v := range nums {
		if _, ok := hashSet[v]; ok {
			return v
		} else {
			hashSet[v] = 1
		}
	}
	return 0
}