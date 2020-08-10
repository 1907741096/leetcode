package main

import "strconv"

func movingCount(m int, n int, k int) int {
	arr := make([][]bool, m)
	for i := 0; i < m; i ++ {
		arr[i] = make([]bool, n)
	}
	num := 0
	queue := [][]int{{0,0}}

	for len(queue) != 0 {
		point := queue[0]
		queue = queue[1:]
		if getSum(point[0], point[1]) <= k && !arr[point[0]][point[1]] {
			num ++
			if point[0] > 0 && !arr[point[0] - 1][point[1]] {
				queue = append(queue, []int{point[0] - 1, point[1]})
			}
			if point[1] > 0 && !arr[point[0]][point[1] - 1] {
				queue = append(queue, []int{point[0], point[1] - 1})
			}
			if point[0] < m - 1 && !arr[point[0] + 1][point[1]] {
				queue = append(queue, []int{point[0] + 1, point[1]})
			}
			if point[1] < n - 1 && !arr[point[0]][point[1] + 1] {
				queue = append(queue, []int{point[0], point[1] + 1})
			}
		}
		arr[point[0]][point[1]] = true
	}
	return num
}

func getSum(x, y int) int {
	sum := 0
	str1 := strconv.Itoa(x)
	for i := 0; i < len(str1); i ++ {
		sum += int(str1[i] - '0')
	}
	str2 := strconv.Itoa(y)
	for i := 0; i < len(str2); i ++ {
		sum += int(str2[i] - '0')
	}
	return sum
}