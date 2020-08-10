package main

func calculateMinimumHP(dungeon [][]int) int {
	minNum := 0
	arr := make([][]int, len(dungeon))
	for i,_ := range dungeon {
		arr[i] = make([]int, len(dungeon[i]))
		for j,_ := range dungeon[i] {
			if i == 0 && j == 0 {
				arr[i][j] = dungeon[i][j]
			} else if i == 0 {
				arr[i][j] = arr[i][j - 1] + dungeon[i][j]
			}  else if j == 0 {
				arr[i][j] = arr[i - 1][j] + dungeon[i][j]
			} else {
				arr[i][j] = max(arr[i - 1][j], arr[i][j - 1]) + dungeon[i][j]
			}
			if arr[i][j] < minNum {
				minNum = arr[i][j]
			}
		}
	}
	return - minNum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}