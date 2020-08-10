package main

func kthSmallest(matrix [][]int, k int) int {
	index := make([]int, len(matrix))
	for j := 0; j < len(matrix); j ++ {
		index[j] = 0
	}
	minIndex := 0
	for i := 0; i < k; i ++ {
		min := matrix[len(matrix) - 1][len(matrix) - 1]
		for j := 0; j < len(matrix); j ++ {
			if index[j] < len(matrix) && matrix[j][index[j]] <= min {
				min = matrix[j][index[j]]
				minIndex = j
			}
		}
		index[minIndex] ++
	}
	return matrix[minIndex][index[minIndex] - 1]
}