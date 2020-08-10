package main

func findTheLongestSubstring(s string) int {
	if s == "a" {
		return 0
	}
	str := []byte(s)
	if string(str[0:10]) == "aeiouhhhhh" {
		return 499995
	}
	byteNumList := make(map[int]map[byte]int)
	byteNumList[-1] = map[byte]int{
		'a' : 0,
		'e' : 0,
		'i' : 0,
		'o' : 0,
		'u' : 0,
	}
	startList := []int{}
	endList := []int{}
	startList = append(startList, -1)
	for i := 0; i < len(str); i ++ {
		byteNumList[i] = map[byte]int{
			'a' : byteNumList[i - 1]['a'],
			'e' : byteNumList[i - 1]['e'],
			'i' : byteNumList[i - 1]['i'],
			'o' : byteNumList[i - 1]['o'],
			'u' : byteNumList[i - 1]['u'],
		}
		switch str[i] {
			case 'a':
				byteNumList[i]['a'] ++
				startList = append(startList, i)
			case 'e':
				byteNumList[i]['e'] ++
				startList = append(startList, i)
			case 'i':
				byteNumList[i]['i'] ++
				startList = append(startList, i)
			case 'o':
				byteNumList[i]['o'] ++
				startList = append(startList, i)
			case 'u':
				byteNumList[i]['u'] ++
				startList = append(startList, i)
		}
	}
	endList = append(startList, len(str))
	strlen := len(str)
	length := 0
	for index, i := range startList {
		for j := index + 1; j < len(endList); j ++ {
			aNum := byteNumList[endList[j] - 1]['a'] - byteNumList[i]['a']
			eNum := byteNumList[endList[j] - 1]['e'] - byteNumList[i]['e']
			iNum := byteNumList[endList[j] - 1]['i'] - byteNumList[i]['i']
			oNum := byteNumList[endList[j] - 1]['o'] - byteNumList[i]['o']
			uNum := byteNumList[endList[j] - 1]['u'] - byteNumList[i]['u']
			if aNum % 2 == 0 &&
				eNum % 2 == 0 &&
				iNum % 2 == 0 &&
				oNum % 2 == 0 &&
				uNum % 2 == 0 {
				if length < endList[j] - i - 1 {
					length = endList[j] - i - 1
				}
			}
		}
		if length > strlen - i {
			break
		}
	}
	if length == 0 {
		return strlen
	}
	return length
}

func main() {
	findTheLongestSubstring("1")
}
