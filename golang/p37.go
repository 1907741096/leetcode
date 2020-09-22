package main

import "fmt"

var dir = []int{-1, 0, 1}

func solveSudoku(board [][]byte)  {
	getSduku(0, 0, &board)
}

func getSduku(x, y int, board *[][]byte) bool {
	if x == 0 && y == 9 {
		return true
	}
	//fmt.Printf("%d,%d,%c\n", x, y , (*board)[x][y])
	if (*board)[x][y] != '.' {
		if x == 8 {
			return getSduku(0, y + 1, board)
		} else {
			return getSduku(x + 1, y, board)
		}
	} else {
		numFlag := make([]bool, 9)
		for i := 0; i < 9; i ++ {
			if (*board)[x][i] != '.' {
				numFlag[(*board)[x][i] - '1'] = true
			}
			if (*board)[i][y] != '.' {
				numFlag[(*board)[i][y] - '1'] = true
			}
		}
		row, col := x / 3, y / 3
		for i := 0; i < 3; i ++ {
			for j := 0; j < 3; j ++ {
				if (*board)[row * 3 + 1 + dir[i]][col * 3 + 1 + dir[j]] != '.' {
					numFlag[(*board)[row * 3 + 1 + dir[i]][col * 3 + 1 + dir[j]] - '1'] = true
				}
			}
		}
		//fmt.Println(numFlag)
		//time.Sleep(time.Second)

		for i, v := range numFlag {
			if !v {
				flag := false
				(*board)[x][y] = byte(i + '1')
				if x == 8 {
					flag = getSduku(0, y + 1, board)
				} else {
					flag = getSduku(x + 1, y, board)
				}
				if flag {
					return true
				} else {
					(*board)[x][y] = '.'
				}
			}
		}
	}

	return false
}

func main() {
	board := [][]byte{
		{'5','3','.','.','7','.','.','.','.'},
		{'6','.','.','1','9','5','.','.','.'},
		{'.','9','8','.','.','.','.','6','.'},
		{'8','.','.','.','6','.','.','.','3'},
		{'4','.','.','8','.','3','.','.','1'},
		{'7','.','.','.','2','.','.','.','6'},
		{'.','6','.','.','.','.','2','8','.'},
		{'.','.','.','4','1','9','.','.','5'},
		{'.','.','.','.','8','.','.','7','9'},
	}
	solveSudoku(board)
	fmt.Println(board)
}

func solveSudoku2(board [][]byte) {
	var line, column [9][9]bool
	var block [3][3][9]bool

	for i, row := range board {
		for j, b := range row {
			if b != '.' {
				digit := b - '1'
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
			}
		}
	}

	var dfs func(int, int) bool
	dfs = func(i, j int) bool {
		if i == 0 && j == 9 {
			return true
		}
		if board[i][j] == '.' {
			for digit := 0; digit < 9; digit++ {
				if !line[i][digit] && !column[j][digit] && !block[i/3][j/3][digit] {
					line[i][digit] = true
					column[j][digit] = true
					block[i/3][j/3][digit] = true
					board[i][j] = byte('1' + digit)
					flag := false

					if i == 8 {
						flag = dfs(0, j+1)
					} else {
						flag = dfs(i+1, j)
					}
					if flag {
						return true
					} else {
						board[i][j] = '.'
					}

					line[i][digit] = false
					column[j][digit] = false
					block[i/3][j/3][digit] = false
				}
			}
		} else {
			if i == 8 {
				return dfs(0, j+1)
			} else {
				return dfs(i+1, j)
			}
		}
		return false
	}
	dfs(0, 0)
}
