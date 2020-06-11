package main

// 哈希表
func dailyTemperatures(T []int) []int {
	length := len(T)
	arr := make([]int, length)
	tHash := make(map[int][]int)
	for i := 0; i < length; i ++ {
		for t, indexArr := range tHash {
			if t < T[i] {
				for j := 0; j < len(indexArr); j ++ {
					if arr[indexArr[j]] == 0 {
						arr[indexArr[j]] = i - indexArr[j]
						if j >= len(tHash[t]) - 1 {
							tHash[t] = append(tHash[t][0: j], []int{}...)
						} else {
							tHash[t] = append(tHash[t][0: j], tHash[t][j + 1: len(tHash[t])]...)
						}
						j --
					}
				}
			}
		}
		if _, ok := tHash[T[i]]; !ok {
			tHash[T[i]] = []int{}
		}
		tHash[T[i]] = append(tHash[T[i]], i)
	}
	return arr
}

// 栈
func dailyTemperatures2(T []int) []int {
	length := len(T)
	arr := make([]int, length)
	stack := []int{}
	for i := 0; i < length; i ++ {
		for len(stack) != 0 && T[i] > T[stack[len(stack) - 1]]  {
			arr[stack[len(stack) - 1]] = i - stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, i)
	}
	return arr
}