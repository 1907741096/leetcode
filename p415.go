package main

import (
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	var num string
	for i, j := len(num1) - 1, len(num2) - 1; i >= 0 || j >= 0; {
		var n int
		if i < 0 {
			n = int(num2[j] - '0')
			j --
		} else if j < 0 {
			n = int(num1[i] - '0')
			i --
		} else {
			n = int(num1[i] - '0') + int(num2[j] - '0')
			i --
			j --
		}
		if num == "" {
			if n >= 10 {
				num = strconv.Itoa(n)
			} else {
				num = "0" + strconv.Itoa(n)
			}
		} else {
			s := n + int(num[0] - '0')
			if s >= 10 {
				num = strconv.Itoa(s) + num[1:]
			} else {
				num = "0" + strconv.Itoa(s) + num[1:]
			}
		}
	}
	if num[0] == '0' {
		num = num[1:]
	}
	return num
}