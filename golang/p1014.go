package main

func maxScoreSightseeingPair(A []int) int {
	index := A[0]
	max := 0
	for i := 1; i < len(A); i ++ {
		if max < index + A[i] - i {
			max = index + A[i] - i
		}
		if index < A[i] + i {
			index = A[i] + i
		}
	}
	return max
}