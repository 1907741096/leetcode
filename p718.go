package main

func findLength(A []int, B []int) int {
	aLength, bLength := len(A), len(B)
	dp := make([][]int, aLength)
	maxLength := 0
	for i := 0; i < aLength; i ++ {
		dp[i] = make([]int, bLength)
		for j := 0; j < bLength; j ++ {
			if A[i] == B[j] {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i - 1][j - 1] + 1
				}
			} else {
				dp[i][j] = 0
			}
			if dp[i][j] > maxLength {
				maxLength = dp[i][j]
			}
		}
	}
	return maxLength
}