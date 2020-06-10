package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n > 0 {
		return quickMi(x, n)
	} else {
		return quickMi(1 / x, -n)
	}
}

func quickMi(x float64, n int) float64 {
	if n <= 1 {
		return x
	}
	if n % 2 == 0 {
		return quickMi(x * x, n / 2)
	} else {
		return quickMi(x * x, n / 2) * x
	}
}

func main() {
	fmt.Println(myPow(2, 6))
}