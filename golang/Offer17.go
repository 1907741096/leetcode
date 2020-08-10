package main

import "strconv"

func printNumbers(n int) []int {
	str := ""
	for i := 0; i < n; i ++ {
		str += "9"
	}
	num, _ := strconv.Atoi(str)
	arr := []int{}
	for i := 1; i <= num; i ++ {
		arr = append(arr, i)
	}
	return arr
}