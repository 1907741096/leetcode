package main

func longestPalindrome(s string) string {
	str := []byte(s)
	length := len(str)
	dp := make([][]int, length)
	maxLen := 0
	maxStr := ""
	for i := 0; i < length; i ++ {
		dp[i] = make([]int, length)
	}
	// start 开始位置
	// add 偏移量
	// dp[开始位置][偏移量是否回文]
	for add := 0; add < length; add ++ {
		for start := 0; start < length - add; start ++ {
			if add == 0 {
				dp[start][add] = 1
			} else if add == 1 {
				if str[start] == str[start + add] {
					dp[start][add] = 1
				} else {
					dp[start][add] = 0
				}
			} else {
				if str[start] == str[start + add] && dp[start + 1][add - 2] == 1 {
					dp[start][add] = 1
				} else {
					dp[start][add] = 0
				}
			}

			if dp[start][add] == 1 && add + 1 > maxLen {
				maxLen = add + 1
				maxStr = s[start: start + add + 1]
			}
		}
	}
	return maxStr
}