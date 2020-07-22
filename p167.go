package main

func twoSum(numbers []int, target int) []int {
	hashSet := make(map[int]int)
	for i, v := range numbers {
		if index, ok := hashSet[target - v]; ok {
			return []int{index + 1, i + 1}
		}
		hashSet[v] = i
	}
	return []int{}
}