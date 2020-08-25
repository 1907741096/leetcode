package main

import "fmt"

func repeatedSubstringPattern(s string) bool {
	length := len(s)
	for i := 1; i <= length / 2; i ++ {
		if length % i != 0 {
			continue
		}
		flag := true
		for j := i; j < length; j += i {
			if s[0:i] != s[j:j+i] {
				flag = false
				break
			}
		}
		if flag {
			return true
		}
	}
	return false
}