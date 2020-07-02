package main

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	n3 := n / 3
	m := n % 3
	if m == 1 {
		m = 4
		n3 --
	}
	if m == 0 {
		m = 1
	}
	return (getThree(n3) * m) % 1000000007
}

func getThree(n int) int {
	res := 1
	for i := 0; i < n; i ++ {
		res = (res * 3) % 1000000007
	}
	return res
}