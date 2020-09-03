package main

var solutions [][]string

func solveNQueens(n int) [][]string {
	solutions = [][]string{}
	queenMap := make([][]bool, n)
	indexList := make([]int, n)
	for i := 0; i < n; i ++ {
		queenMap[i] = make([]bool, n)
	}
	backtrack(0, n, queenMap, indexList)
	return solutions
}

func backtrack(row, n int, queenMap [][]bool, indexList []int) {
	if n == row {
		solutions = append(solutions, generateBoard(indexList))
		return
	}

	nowQueenMap := make([][]bool, n)
	for j := 0; j < n; j ++ {
		nowQueenMap[j] = make([]bool, n)
	}
	for i := 0; i < n; i ++ {
		if queenMap[row][i] {
			continue
		}
		nowQueenMap[row][i] = true
		for left, right, height := i - 1, i + 1, row + 1; height < n; height ++ {
			copy(nowQueenMap[height], queenMap[height])
			if left >= 0 {
				nowQueenMap[height][left] = true
				left --
			}
			if right < n {
				nowQueenMap[height][right] = true
				right ++
			}
			nowQueenMap[height][i] = true
		}
		indexList[row] = i
		backtrack(row + 1, n, nowQueenMap, indexList)
	}
}

func generateBoard(queens []int) []string {
	board := []string{}
	for i := 0; i < len(queens); i++ {
		row := make([]byte, len(queens))
		for j := 0; j < len(queens); j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}