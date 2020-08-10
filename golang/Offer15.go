package main

func hammingWeight(num uint32) int {
	n := 0
	for num > 0 {
		if num % 2 == 1 {
			n ++
		}
		num /= 2
	}
	return n
}