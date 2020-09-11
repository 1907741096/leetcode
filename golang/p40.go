package main

import (
	"sort"
	"strconv"
)

var resArr [][]int

var strSet map[string]bool

var ca []int

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	resArr = [][]int{}
	strSet = make(map[string]bool)
	ca = candidates
	getCombine2(0, len(ca), target, []int{}, "")
	return resArr
}

func getCombine2(i, length int, target int, list []int, s string) {
	for j := i; j < length; j ++ {
		if target - ca[j] < 0 {
			continue
		} else if target - ca[j] == 0 {
			if !strSet[s + "," + strconv.Itoa(ca[j])] {
				resArr = append(resArr, append(list, ca[j]))
				strSet[s + "," + strconv.Itoa(ca[j])] = true
			}
		} else {
			newList := make([]int, len(list))
			copy(newList, list)
			newList = append(newList, ca[j])
			getCombine2(j + 1, length, target - ca[j], newList, s + "," + strconv.Itoa(ca[j]))
		}
	}
}