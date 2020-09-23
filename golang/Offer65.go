package main

func add(a int, b int) int {
	for b != 0 {
		c := (a & b) << 1 // 算进位
		a ^= b // 算不包括进位的保留位
		b = c
	}
	return a
}