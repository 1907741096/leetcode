package main

// 哈希表
func dailyTemperatures(T []int) []int {
	length := len(T)
	arr := make([]int, length)
	tHash := make(map[int][]int)
	for i := 0; i < length; i ++ {
		for t, _ := range tHash {
			if t < T[i] {
				for len(tHash[t]) != 0 {
					arr[tHash[t][0]] = i - tHash[t][0]
					if len(tHash[t]) == 1 {
						tHash[t] = []int{}
					} else {
						tHash[t] = tHash[t][1:]
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