package main

var resArr [][]int

var ca []int

func combinationSum(candidates []int, target int) [][]int {
	resArr = [][]int{}
	ca = candidates
	length := len(ca)
	for i := 0; i < length; i ++ {
		getCombine(i, length, target - ca[i], []int{ca[i]})
	}
	return resArr
}

func getCombine(i, length int, target int, list []int) {
	if target < 0 || i >= length {
		return
	}
	if target == 0 {
		resArr = append(resArr, list)
	} else {
		for j := i; j < length; j ++ {
			newList := make([]int, len(list))
			copy(newList, list)
			newList = append(newList, ca[j])
			getCombine(j, length, target - ca[j], newList)
		}
	}
}