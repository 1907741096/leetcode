package main

import "math"

func integerBreak(n int) int {
	a := n / 3
	if a == 0 {
		return 1
	}
	if n == 3 {
		return 2
	}
	b := n % 3
	if b == 0 {
		return int(math.Pow(float64(3), float64(a)))
	} else if b == 1 {
		return 4 * int(math.Pow(float64(3), float64(a - 1)))
	} else {
		return 2 * int(math.Pow(float64(3), float64(a)))
	}
}