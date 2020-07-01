package main

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n < 3 {
		return 1
	}
	a := 1
	b := 1
	c := 1
	for i := 2; i < n; i ++ {
		c = a + b
		a = b
		b = c % 1000000007
	}
	return c % 1000000007
}