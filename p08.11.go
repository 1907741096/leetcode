package main

func waysToChange(n int) int {
	coins := []int{25, 10, 5, 1}
	dp := make([]int, n + 1)
	dp[0] = 1
	for j := 0; j < len(coins); j ++ {
		for i := 1; i <= n; i ++ {
			if i - coins[j] >= 0 {
				dp[i] = (dp[i] + dp[i - coins[j]]) % 1000000007
			}
		}
	}
	return dp[n]
}