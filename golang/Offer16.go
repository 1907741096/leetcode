package main

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	var res = x
	var now = 1.0
	for n > 1 {
		if n % 2 == 1 {
			now *= res
		}
		res *= res
		n /= 2
	}
	return res * now
}