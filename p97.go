package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	dp := make([]bool, len(s3) + 1)
	dp[0] = true
	for i := 1; i <= len(s3); i ++ {
		for j, k := len(s1) - 1, i; j >= 0; j -- {
			if k > 0 && s3[k - 1] == s1[j] {
				k --
				if j == 0 {
					dp[i] = dp[k]
				}
			} else {
				break
			}
		}
		for j, k := len(s2) - 1, i; j >= 0; j -- {
			if k > 0 && s3[k - 1] == s2[j] {
				k --
				if j == 0 && dp[i] == false {
					dp[i] = dp[k]
				}
			} else {
				break
			}
		}
	}
	return dp[len(s3)]
}