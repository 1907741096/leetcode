package main

func lastRemaining(n int, m int) int {
	arr := []int{}
	for i := 0; i < n; i ++ {
		arr = append(arr, i)
	}
	index := -1
	for len(arr) != 1 {
		index += m
		if index >= len(arr) {
			index %= len(arr)
		}
		arr = append(arr[:index], arr[index + 1:]...)
		index --
	}
	return arr[0]
}