package main

import "strings"

func isPalindrome(s string) bool {
	i, j := 0, len(s) - 1
	s = strings.ToLower(s)
	for i < j {
		if !(s[i] >= '0' && s[i] <= '9' || s[i] <= 'z' && s[i] >= 'a') {
			i ++
			continue
		}
		if !(s[j] >= '0' && s[j] <= '9' || s[j] <= 'z' && s[j] >= 'a') {
			j --
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i ++
		j --
	}
	return true
}
