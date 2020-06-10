package main

var arrNum = []int{}

func isPalindrome(x int) bool {
	arrNum = []int{}
	if x < 0 {
		return false
	}
	getOne(x)
	length := len(arrNum)
	for i := 0; i < length / 2; i ++ {
		if arrNum[i] != arrNum[length - i - 1] {
			return false
		}
	}
	return true
}

func getOne(x int) {
	if x > 9 {
		arrNum = append(arrNum, x % 10)
		getOne(x / 10)
	} else {
		arrNum = append(arrNum, x % 10)
	}
}