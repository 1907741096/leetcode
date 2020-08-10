package main

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}
	str := strs[0]
	if str == "" {
		return ""
	}
	for i := 1; i < length; i ++ {
		strLength := len(strs[i])
		if strLength == 0 {
			return ""
		}
		for j := 0; j < strLength; j ++ {
			if j == len(str) {
				break
			}
			if str[j] != strs[i][j] {
				if j == 0 {
					return ""
				}
				str = str[0: j]
				break
			}
			if j == strLength - 1 {
				str = str[0: strLength]
			}
		}
	}
	return str
}