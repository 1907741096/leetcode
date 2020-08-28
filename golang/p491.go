package main

import "strconv"

type t struct {
	value []int
	str string
}

func findSubsequences(nums []int) [][]int {
	arr := []t{}
	for i := 0; i < len(nums); i ++ {
		length := len(arr)
		for j := 0; j < length; j ++ {
			if arr[j].value[len(arr[j].value) - 1] <= nums[i] {
				newArr := make([]int, len(arr[j].value))
				copy(newArr, arr[j].value)
				newT := t{
					value: append(newArr, nums[i]),
					str:   arr[j].str + strconv.Itoa(nums[i]) + ",",
				}
				arr = append(arr, newT)
			} else {
				break
			}
		}
		arr = append(arr, t{
			value: []int{nums[i]},
			str:   strconv.Itoa(nums[i]),
		})
	}
	res := [][]int{}
	hashSet := map[string]bool{}
	for _, a := range arr {
		if len(a.value) != 1 && !hashSet[a.str] {
			res = append(res, a.value)
			hashSet[a.str] = true
		}
	}
	return res
}