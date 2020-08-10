package main

type firstChar struct {
	flag bool
	index int
}

func firstUniqChar(s string) byte {
	hashSet := make(map[byte]firstChar)
	for i, v := range []byte(s) {
		if _, ok := hashSet[v]; !ok {
			hashSet[v] = firstChar{
				flag: true,
				index: i,
			}
		} else {
			hashSet[v] = firstChar{
				flag: false,
			}
		}
	}

	min := -1
	for _, v := range hashSet {
		if v.flag {
			if min == -1 || min > v.index {
				min = v.index
			}
		}
	}
	if min == -1 {
		return ' '
	}
	return s[min]
}