package main

import "strconv"

func isPathCrossing(path string) bool {
	x, y := 0, 0
	hastSet := make(map[string]bool)
	hastSet["0,0"] = true
	for _, v := range path {
		if v == 'N' {
			y ++
		} else if v == 'S' {
			y --
		} else if v == 'E' {
			x ++
		} else {
			x --
		}
		if hastSet[strconv.Itoa(x) + "," + strconv.Itoa(y)] {
			return true
		} else {
			hastSet[strconv.Itoa(x) + "," + strconv.Itoa(y)] = true
		}
	}
	return false
}