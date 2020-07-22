package main

func minArray(numbers []int) int {
	left, right := 0, len(numbers) - 1
	minNum := numbers[0]
	for left <= right {
		mid := left + (right - left) / 2
		if minNum > numbers[mid] {
			minNum = numbers[mid]
		}
		if minNum > numbers[left] {
			minNum = numbers[left]
		}
		if numbers[mid] == numbers[left] {
			left ++
		} else if numbers[mid] > numbers[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return minNum
}