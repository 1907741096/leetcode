package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[len(obstacleGrid) - 1][len(obstacleGrid[0]) - 1] == 1 {
		return 0
	}
	arr := make([][]int, len(obstacleGrid))
	for i,_ := range obstacleGrid {
		arr[i] = make([]int, len(obstacleGrid[i]))
		for j,_ := range obstacleGrid[i] {
			if i == 0 && j == 0 {
				if obstacleGrid[0][0] == 1 {
					arr[0][0] = 0
				} else {
					arr[0][0] = 1
				}
				continue
			}
			top := 0
			left := 0
			if j - 1 >= 0 && obstacleGrid[i][j - 1] != 1 {
				top = arr[i][j - 1]
			}
			if i - 1 >= 0 && obstacleGrid[i - 1][j] != 1 {
				left = arr[i - 1][j]
			}
			arr[i][j] = top + left
		}
	}
	return arr[len(obstacleGrid) - 1][len(obstacleGrid[0]) - 1]
}
