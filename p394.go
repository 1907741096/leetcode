package main

import (
	"strconv"
)

func decodeString(s string) string {
	str, _ := getStr([]byte(s))
	return str
}

func getStr(arr []byte) (string, int) {
	num := 1
	sumStr := ""
	index := -1
	for i := 0; i < len(arr); i ++ {
		if arr[i] >= '0' && arr[i] <= '9'{
			if index == -1 {
				index = i
			}
		} else if arr[i] == '[' {
			if i != 0 {
		 		num, _ = strconv.Atoi(string(arr[index:i]))
				index = -1
			}
			childStr, n := getStr(arr[i + 1:])
			for j := 0; j < num; j ++ {
				sumStr += childStr
			}
			i += (n + 1)
		} else if arr[i] == ']' {
			return sumStr, i
		} else {
			sumStr += string(arr[i])
		}
	}
	return sumStr, len(arr)
}
