package main

var resArr [][]int

func combinationSum3(k int, n int) [][]int {
	resArr = [][]int{}
	getCombine3(1, k, n, []int{})
	return resArr
}

func getCombine3(i, k int, target int, list []int) {
	for j := i; j < 10; j ++ {
		if target - j < 0 {
			continue
		} else if target - j == 0 {
			if len(list) == k - 1 {
				resArr = append(resArr, append(list, j))
			}
			break
		} else {
			newList := make([]int, len(list))
			copy(newList, list)
			newList = append(newList, j)
			getCombine3(j + 1, k, target - j, newList)
		}
	}
}