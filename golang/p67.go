package main

func addBinary(a string, b string) string {
	str := ""
	flag := false
	i, j := len(a) - 1, len(b) - 1
	for {
		if i < 0 && j < 0 {
			if flag {
				str = "1" + str
			}
			break
		}
		strOne := ""
		if i < 0 {
			if flag {
				if b[j] == '1' {
					strOne = "0"
					flag = true
				} else {
					strOne = "1"
					flag = false
				}
				str = strOne + str
				j --
			} else {
				str = b[0: j + 1] + str
				break
			}
		} else if j < 0 {
			if flag {
				if a[i] == '1' {
					strOne = "0"
					flag = true
				} else {
					strOne = "1"
					flag = false
				}
				str = strOne + str
				i --
			} else {
				str = a[0: i + 1] + str
				break
			}
		} else {
			if flag {
				if a[i] == '1' && b[j] == '1' {
					flag = true
					strOne = "1"
				} else if a[i] == '0' && b[j] == '0' {
					flag = false
					strOne = "1"
				} else {
					flag = true
					strOne = "0"
				}
			} else {
				if a[i] == '1' && b[j] == '1' {
					flag = true
					strOne = "0"
				} else if a[i] == '0' && b[j] == '0' {
					flag = false
					strOne = "0"
				} else {
					flag = false
					strOne = "1"
				}
			}
			str = strOne + str
			i --
			j --
		}
	}
	return str
}