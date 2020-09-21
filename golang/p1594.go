package main

func maxProductPath(grid [][]int) int {
	maxNum := -1
	var f func(x, y, num int)
	f = func(x, y ,num int) {
		if num == 0 || x == len(grid) - 1 && y == len(grid[0]) - 1 {
			if num > maxNum {
				maxNum = num
			}
			return
		}
		if x < len(grid) - 1 {
			f(x + 1, y, num * grid[x + 1][y])
		}
		if y < len(grid[0]) - 1 {
			f(x, y + 1, num * grid[x][y + 1])
		}
	}
	f(0, 0, grid[0][0])
	if maxNum == -1 {
		return maxNum
	} else {
		return maxNum % 1000000007
	}
}