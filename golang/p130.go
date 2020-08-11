package main

var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

var flagMap [][]bool
var boardMap *[][]byte

func solve(board [][]byte)  {
	flagMap = make([][]bool, len(board))
	boardMap = &board
	for i := 0; i < len(board); i ++ {
		flagMap[i] = make([]bool, len(board[i]))
	}
	for i := 0; i < len(board); i ++ {
		for j := 0; j < len(board[i]); j ++ {
			if i == 0 || j == 0 || i == len(board) - 1 || j == len(board[i]) - 1 {
				dfs(i, j)
			}
		}
	}
	for i := 0; i < len(board); i ++ {
		for j := 0; j < len(board[i]); j ++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(i, j int) {
	if flagMap[i][j] {
		return
	}
	flagMap[i][j] = true
	if (*boardMap)[i][j] == 'O' {
		(*boardMap)[i][j] = 'A'
		for _, d := range dir {
			x := i + d[0]
			y := j + d[1]
			if x < 0 || x >= len(*boardMap) || y < 0 || y >= len((*boardMap)[i]) {
				// 超出边界
			} else {
				dfs(x, y)
			}
		}
	}
}