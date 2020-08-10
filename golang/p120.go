package main

import "math"

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	arr := make([]int, len(triangle[len(triangle) - 1]))
	minNum := math.MaxInt32
	for i := 0; i < len(triangle); i ++ {
		for j := len(triangle[i]) - 1; j >= 0; j -- {
			if j == 0 {
				arr[j] = arr[j] + triangle[i][j]
			} else if j == len(triangle[i]) - 1 {
				arr[j] = arr[j - 1] + triangle[i][j]
			} else {
				arr[j] = min(arr[j] + triangle[i][j], arr[j - 1] + triangle[i][j])
			}
			if i == len(triangle) - 1 && minNum > arr[j] {
				minNum = arr[j]
			}
		}
	}
	return minNum
}

