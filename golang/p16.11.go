package main

func divingBoard(shorter int, longer int, k int) []int {
	res := []int{}
	if k == 0 {
		return res
	}
	for i := 0; i <= k; i ++ {
		num := longer * i + shorter * (k - i)
		if len(res) == 0 || res[len(res) - 1] != num {
			res = append(res, longer * i + shorter * (k - i))
		}
	}
	return res
}
