package main

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	min := 0
	max := x
	mid := 0
	for {
		mid = (min + max) / 2
		if mid * mid <= x && (mid + 1) * (mid + 1) > x{
			break
		}
		if mid * mid > x {
			max = mid
		} else {
			min = mid
		}
	}
	return mid
}