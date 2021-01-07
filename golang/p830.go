package main

func largeGroupPositions(s string) [][]int {
	var res [][]int
	pre := 0
	for i := 1; i < len(s); i ++ {
		if s[i] != s[pre] {
			if i - pre >= 3 {
				res = append(res, []int{pre, i - 1})
			}
			pre = i
		}
		if i == len(s) - 1 {
			if i - pre >= 2 {
				res = append(res, []int{pre, i})
			}
		}
	}
	return res
}