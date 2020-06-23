package main

import "fmt"

func trapRainWater(heightMap [][]int) int {
	//heightMap = [][]int{
	//	{5,5,5,1},
	//	{5,1,1,5},
	//	{5,1,5,5},
	//	{5,2,5,8},
	//}
	height := len(heightMap)
	width := len(heightMap[0])
	leftMax := make([][]int, height)
	rightMax := make([][]int, height)
	topMax := make([][]int, height)
	bottomMax := make([][]int, height)
	for i := 0; i < height; i ++ {
		leftMax[i] = make([]int, width)
		topMax[i] = make([]int, width)
		rightMax[i] = make([]int, width)
		bottomMax[height - i - 1] = make([]int, width)
		for j := 0; j < width; j ++ {
			if i == 0 {
				if j == 0 {
					leftMax[i][j] = heightMap[i][j]
					rightMax[i][width - j - 1] = heightMap[i][width - j - 1]
				} else {
					leftMax[i][j] = max(leftMax[i][j - 1], heightMap[i][j])
					rightMax[i][width - j - 1] = max(rightMax[i][width - j], heightMap[i][j])
				}
				topMax[i][j] = heightMap[i][j]
				bottomMax[height - i - 1][j] = heightMap[height - i - 1][j]
			} else {
				if j == 0 {
					leftMax[i][j] = heightMap[i][j]
					rightMax[i][width - j - 1] = heightMap[i][width - j - 1]
				} else {
					leftMax[i][j] = max(leftMax[i][j - 1], heightMap[i][j])
					rightMax[i][width - j - 1] = max(rightMax[i][width - j], heightMap[i][j])
				}
				topMax[i][j] = max(topMax[i - 1][j], heightMap[i][j])
				bottomMax[height - i - 1][j] = max(bottomMax[height - i][j], heightMap[height - i - 1][j])
			}
		}
	}
	fmt.Println(leftMax)
	fmt.Println(rightMax)
	fmt.Println(topMax)
	fmt.Println(bottomMax)
	sum := 0
	for i := 1; i < height - 1; i ++ {
		for j := 1; j < width - 1; j ++ {
			minHeight := min(min(leftMax[i][j - 1], rightMax[i][j + 1]), min(topMax[i - 1][j], bottomMax[i + 1][j]))
			if minHeight > heightMap[i][j] {
				sum += minHeight - heightMap[i][j]
			}
		}
	}
	return sum
}