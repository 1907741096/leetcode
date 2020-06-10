package main

import (
	"math"
)

func maximalSquare(matrix [][]byte) int {
	max := 0
	numArr := make([][]int, len(matrix))
	for i, _ := range numArr {
		numArr[i] = make([]int, len(matrix[i]))
	}
	for x, value := range matrix {
		for y, i := range value {
			if i == '1' {
				var topNum, leftNum, leftTopNum int
				if x - 1 < 0 {
					topNum = 0
				} else {
					topNum = numArr[x - 1][y]
				}
				if y - 1 < 0 {
					leftNum = 0
				} else {
					leftNum = numArr[x][y - 1]
				}
				if y > 0 && x > 0 {
					leftTopNum = numArr[x - 1][y - 1]
				} else {
					leftTopNum = 0
				}
				if x == 0 || y == 0 {
					numArr[x][y] = int(math.Max(float64(leftNum), float64(topNum)))
					if numArr[x][y] == 0 {
						numArr[x][y] = 1
					}
				} else {
					if leftNum > 0 && topNum > 0 && leftTopNum > 0 {
						numArr[x][y] = int(math.Min(math.Min(float64(leftNum), float64(topNum)), float64(leftTopNum))) + 1
					} else {
						numArr[x][y] = 1
					}
				}
			} else {
				numArr[x][y] = 0
			}
			if numArr[x][y] > max {
				max = numArr[x][y]
			}
		}
	}
	return max * max
}