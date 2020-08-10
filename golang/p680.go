package main

func validPalindrome(s string) bool {
	isOne := false
	str := []byte(s)
	length := len(str) - 1
	i := 0
	j := length
	for i < j {
		if str[i] == str[j] {
			i ++
			j --
		} else {
			if isOne {
				return false
			} else {
				isOne = true
				if str[i + 1] == str[j] {
					if length > 2 && str[i + 2] == str[j - 1] {
						i ++
					} else {
						j --
					}
				} else {
					j --
				}
			}
		}
	}
	return true
}