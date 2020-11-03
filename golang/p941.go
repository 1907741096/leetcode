package main

func validMountainArray(A []int) bool {
	left := 0
	right := len(A) - 1
	if len(A) < 3 {
		return false
	}

	for left + 1 <= right - 1 {
		flag := false
		if A[left] < A[left + 1] {
			left ++
			flag = true
		} else if A[left] == A[left + 1] || left == 0 {
			return false
		}
		if A[right] < A[right - 1] {
			right --
			flag = true
		} else if A[right] == A[right - 1] || right == len(A) - 1 {
			return false
		}
		if !flag {
			return false
		}

	}
	return true
}