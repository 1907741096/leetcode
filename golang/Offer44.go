package main

import (
	"fmt"
	"math"
	"strconv"
)

func findNthDigit(n int) int {
	num := 0
	w := 1
	if n == 0 {
		return 0
	}
	if n > 10 {
		n --
		for n > w * int(math.Pow10(w) * 0.9) {
			n -= w * int(math.Pow10(w) * 0.9)
			num = int(math.Pow10(w))
			w ++
		}
	}
	num += n / w
	str := strconv.Itoa(num)
	res := str[n % w] - '0'
	return int(res)
}

func main() {
	fmt.Println(findNthDigit(13))
}
