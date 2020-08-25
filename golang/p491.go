package main

import "strconv"

func findSubsequences(nums []int) [][]int {
	arr := [][]int{}
	for i := 0; i < len(nums); i ++ {
		length := len(arr)
		for j := 0; j < length; j ++ {
			if arr[j][len(arr[j]) - 1] <= nums[i] {
				newArr := make([]int, len(arr[j]) + 1)
				copy(newArr, arr[j])
				newArr[len(newArr) - 1] = nums[i]
				arr = append(arr, newArr)
			}
		}
		arr = append(arr, []int{nums[i]})
	}
	res := [][]int{}
	hashSet := map[string]bool{}
	for _, a := range arr {
		s := ""
		for _, i := range a {
			s += strconv.Itoa(i) + ","
		}
		if len(a) != 1 && !hashSet[s]{
			res = append(res, a)
		}
		hashSet[s] = true
	}
	return res
}