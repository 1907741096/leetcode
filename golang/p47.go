package main

import "sort"

func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	length := len(nums)
	ans = [][]int{}
	res := []int{}
	flag := make([]bool, length)
	var f func(int)
	f = func(i int) {
		if len(res) == length {
			ans = append(ans, append([]int{}, res...))
		}
		for i, v := range nums {
			if flag[i] || i > 0 && !flag[i - 1] && v == nums[i - 1]{
				continue
			}
			flag[i] = true
			res = append(res, v)
			f(i + 1)
			res = res[:len(res) - 1]
			flag[i] = false
		}
	}
	f(0)
	return
}