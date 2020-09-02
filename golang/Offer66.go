package main

func constructArr(a []int) []int {
	if len(a) == 0 {
		return nil
	}
	leftArr := make([]int, len(a))
	leftArr[0] = a[0]
	rightArr := make([]int, len(a))
	rightArr[len(a) - 1] = a[len(a) - 1]
	for i := 1; i < len(a); i ++ {
		leftArr[i] = leftArr[i - 1] * a[i]
		rightArr[len(a) - i - 1] = rightArr[len(a) - i] * a[len(a) - i - 1]
	}
	b := make([]int, len(a))
	for i := 0; i < len(a); i ++ {
		if i == 0 {
			b[i] = rightArr[1]
		} else if i == len(a) - 1 {
			b[i] = leftArr[len(a) - 2]
		} else {
			b[i] = leftArr[i - 1] * rightArr[i + 1]
		}
	}
	return b
}