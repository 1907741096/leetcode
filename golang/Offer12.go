package main

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i ++ {
		for j := 0; j < len(board[i]); j ++ {
			if board[i][j] == word[0] {
				arr := make([][]bool, len(board))
				for i := 0; i < len(board); i ++ {
					arr[i] = make([]bool, len(board[i]))
				}
				y := i
				x := j
				if Enter(board, word, y, x, 1, &arr) {
					return true
				}
			}
		}
	}
	return false
}

func Enter(board [][]byte, word string, y, x, k int, arr *[][]bool) bool {
	if k == len(word) {
		return true
	}
	(*arr)[y][x] = true
	if y > 0 && !(*arr)[y - 1][x] && board[y - 1][x] == word[k] {
		if Enter(board, word, y - 1, x, k + 1, arr) {
			return true
		}
	}
	if x > 0 && !(*arr)[y][x - 1] && board[y][x - 1] == word[k] {
		if Enter(board, word, y, x - 1, k + 1, arr) {
			return true
		}
	}
	if y < len(board) - 1 && !(*arr)[y + 1][x] && board[y + 1][x] == word[k] {
		if Enter(board, word, y + 1, x, k + 1, arr) {
			return true
		}
	}
	if x < len(board[0]) - 1 && !(*arr)[y][x + 1] && board[y][x + 1] == word[k] {
		if Enter(board, word, y, x + 1, k + 1, arr) {
			return true
		}
	}
	(*arr)[y][x] = false
	return false
}