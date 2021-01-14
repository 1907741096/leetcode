package main

func findCircleNum(isConnected [][]int) int {
	flagArr := make([]bool, len(isConnected))
	num := 0
	var dfs func(n int)
	dfs = func(n int) {
		if !flagArr[n] {
			flagArr[n] = true
		} else {
			return
		}
		for i := 0; i < len(isConnected); i ++ {
			if i != n && isConnected[n][i] == 1 {
				dfs(i)
			}
		}
	}
	for i := 0; i < len(isConnected); i ++ {
		if !flagArr[i] {
			num ++
		}
		dfs(i)
	}
	return num
}