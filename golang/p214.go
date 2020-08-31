package main

func shortestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	reverseS := reverse(s)
	for i := 0; i < len(s) - 1; i ++ {
		if reverseS[i:] == s[:len(s) - i] {
			return reverseS[:i] + s
		}
	}
	return reverseS + s[1:]
}

// 反转字符串
func reverse(word string) string {
	runes := []byte(word)
	for i := 0; i < len(runes) / 2; i ++ {
		runes[i], runes[len(runes) - i - 1] = runes[len(runes) - i - 1], runes[i]
	}
	return string(runes)
}