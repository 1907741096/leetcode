package main

var resArr [][]int

var ca []int

func combinationSum(candidates []int, target int) [][]int {
	resArr = [][]int{}
	ca = candidates
	getCombine(0, len(ca), target, []int{})
	return resArr
}

func getCombine(i, length int, target int, list []int) {
	for j := i; j < length; j ++ {
		if target - ca[j] < 0 {
			continue
		} else if target - ca[j] == 0 {
			resArr = append(resArr, append(list, ca[j]))
		} else {
			newList := make([]int, len(list))
			copy(newList, list)
			newList = append(newList, ca[j])
			getCombine(j, length, target - ca[j], newList)
		}
	}
}