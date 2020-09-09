package main

var resArr [][]int

func combine(n int, k int) [][]int {
	resArr = [][]int{}
	for i := 1; i <= n; i ++ {
		getCombine(i + 1, n, k - 1, []int{i})
	}
	return resArr
}

func getCombine(i, n int, k int, list []int) {
	if i + k > n + 1 {
		return
	}
	if k == 0 {
		resArr = append(resArr, list)
	} else {
		for j := i; j <= n; j ++ {
			newList := make([]int, len(list))
			copy(newList, list)
			newList = append(newList, j)
			getCombine(j + 1, n, k - 1, newList)
		}
	}
}