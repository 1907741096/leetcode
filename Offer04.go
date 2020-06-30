package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0{
		return false
	}
	row, col := 0, len(matrix[0]) - 1
	for col >= 0 && row < len(matrix) {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			row ++
		} else {
			col --
		}
	}
	return false
}