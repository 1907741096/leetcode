package main

import (
	"strconv"
)

func palindromePairs(words []string) [][]int {
	res := [][]int{}
	resHash := make(map[string]bool)
	strHash := make(map[string]int)
	for i, word := range words {
		strHash[reverse(word)] = i
	}
	for i, word := range words {
		if word == "" {
			for j, w := range words {
				if i != j && isPalindrome(w) {
					if _, ok := strHash[strconv.Itoa(i) + "," + strconv.Itoa(j)]; !ok {
						res = append(res, []int{i, j})
						resHash[strconv.Itoa(i) + "," + strconv.Itoa(j)] = true
					}
					if _, ok := strHash[strconv.Itoa(j) + "," + strconv.Itoa(i)]; !ok {
						res = append(res, []int{j, i})
						resHash[strconv.Itoa(j) + "," + strconv.Itoa(i)] = true
					}
				}
			}
		} else {
			for j := 0; j < len(word); j ++ {
				if index, ok := strHash[word[j:]]; ok && i != index && isPalindrome(word[:j]) {
					if _, ok := resHash[strconv.Itoa(index) + "," + strconv.Itoa(i)]; !ok {
						res = append(res, []int{index, i})
						resHash[strconv.Itoa(index) + "," + strconv.Itoa(i)] = true
					}
				}
				if index, ok := strHash[word[:len(word) - j]]; ok && i != index && isPalindrome(word[len(word) - j:]) {
					if _, ok := resHash[strconv.Itoa(i) + "," + strconv.Itoa(index)]; !ok {
						res = append(res, []int{i, index})
						resHash[strconv.Itoa(i) + "," + strconv.Itoa(index)] = true
					}
				}
			}
		}
	}
	return res
}

// 反转字符串
func reverse(word string) string {
	runes := []byte(word)
	for i := 0; i < len(runes) / 2; i ++ {
		runes[i], runes[len(runes) - i - 1] = runes[len(runes) - i - 1], runes[i]
	}
	return string(runes)
}

// 判断回文
func isPalindrome(s string) bool {
	i, j := 0, len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i ++
		j --
	}
	return true
}