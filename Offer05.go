package main

func replaceSpace(s string) string {
	str := ""
	for _, c := range s {
		if c == ' ' {
			str += "%20"
		} else {
			str += string(c)
		}
	}
	return str
}