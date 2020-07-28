package main

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	for i, j := 0, 0; i < len(s) && j < len(t); j ++ {
		if s[i] == t[j] {
			i++
		}
		if i == len(s) {
			return true
		}
	}
	return false
}