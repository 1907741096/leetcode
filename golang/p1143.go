package main

import "math"

func longestCommonSubsequence(text1 string, text2 string) int {
	arr1 := []byte(text1)
	arr2 := []byte(text2)
	arr := make([][]int, len(text1))
	for i := 0; i < len(text1); i ++ {
		arr[i] = make([]int, len(text2))
		for j := 0; j < len(text2); j ++ {
			var vi, vj int
			if j - 1 < 0 {
				vi = 0
			} else {
				vi = arr[i][j - 1]
			}
			if i - 1 < 0 {
				vj = 0
			} else {
				vj = arr[i - 1][j]
			}
			if arr1[i] == arr2[j] {
				arr[i][j] = vj+ 1
			} else {
				arr[i][j] = int(math.Max(float64(vi), float64(vj)))
			}
		}
	}
	return arr[len(text1) - 1][len(text2) - 1]
}