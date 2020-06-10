package main

var arr = map[int][]int{
	0: {0,1},
	1: {1,0},
	2: {0,-1},
	3: {-1,0},
}

func spiralOrder(matrix [][]int) []int {
	res := []int{}
	if len(matrix) == 0 {
		return res
	}
	flag := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i ++ {
		flag[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j ++ {
			flag[i][j] = 0
		}
	}
	num := 0
	for x, y := 0, -1; ; {
		// 需要转向
		if x + arr[num % 4][0] < 0 ||
			x + arr[num % 4][0] > len(flag) - 1 ||
			y + arr[num % 4][1] < 0 ||
			y + arr[num % 4][1] > len(flag[x]) - 1 ||
			flag[x + arr[num % 4][0]][y + arr[num % 4][1]] == 1 {
			num ++
			x += arr[num % 4][0]
			y += arr[num % 4][1]
			if x < 0 ||
				x > len(flag) - 1 ||
				y < 0 ||
				y > len(flag[x]) - 1 ||
				flag[x][y] == 1 {
				break
			}
		} else {
			x += arr[num % 4][0]
			y += arr[num % 4][1]
		}
		res = append(res, matrix[x][y])
		flag[x][y] = 1
	}
	return res
}