package main

import "math"

func minCostConnectPoints(points [][]int) int {
	// 普里姆算法
	mapSet := make(map[int]bool)
	mapSet[0] = true
	res := 0
	for length := len(points) - 1; length > 0; length -- {
		minIndex := 0
		minLength := -1
		for i, _ := range mapSet {
			for j := 1; j < len(points); j ++ {
				if mapSet[j] {
					continue
				}
				if minLength == -1 || int(math.Abs(float64(points[i][0] - points[j][0])) + math.Abs(float64(points[i][1] - points[j][1]))) < minLength {
					minLength = int(math.Abs(float64(points[i][0] - points[j][0])) + math.Abs(float64(points[i][1] - points[j][1])))
					minIndex = j
				}
			}
		}
		res += minLength
		mapSet[minIndex] = true
	}
	return res
}