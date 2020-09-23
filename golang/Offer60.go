package main

import "fmt"

func twoSum(n int) []float64 {
	n = 2
	arr := make([]int, n * 6)
	for i := 0; i < 6; i ++ {
		arr[i] = 1
	}
	res := make([]float64, n * 6 - n + 1)
	for i := 1; i < n; i ++ {
		for j := (i + 1) * 6 - 1; j >= i; j -- {
			arr[j] = 0
			for k := 6; k > 0; k -- {
				if j - k >= 0 {
					arr[j] += arr[j - k]
				}
			}
		}
	}
	sum := 0
	for i := n - 1; i < n * 6; i ++ {
		sum += arr[i]
	}

	for i, j := n - 1, 0; i < n * 6; i ++ {
		res[j] = float64(arr[i]) / float64(sum)
		j ++
	}
	return res
}