package main

import "strconv"

func translateNum(num int) int {
	str := []byte(strconv.Itoa(num))
	length := len(str)
	arr := make([]int, length)
	for i := 0; i < length; i ++ {
		if i == 0 {
			arr[i] = 1
		} else if i == 1 {
			if str[i - 1] == '1' || str[i - 1] == '2' && str[i] < '6' {
				arr[i] = 2
			} else {
				arr[i] = 1
			}
		} else {
			if str[i - 1] == '1' || str[i - 1] == '2' && str[i] < '6' {
				arr[i] = arr[i - 2] + arr[i - 1]
			} else {
				arr[i] = arr[i - 1]
			}
		}
	}
	return arr[length - 1]
}