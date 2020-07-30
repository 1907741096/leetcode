package main

type point struct {
	x, y int
}

var dir = []int{
	-1, 0,
	0, 1,
	1, 0,
	0, -1,
}

func minimalSteps(maze []string) int {
	mazes := [][]byte{}
	mList := []point{}
	for i, str := range maze {
		mazes[i] = []byte{}
		for j, c := range str {
			if c == 'M' {
				mList = append(mList, point{
					x: i,
					y: j,
				})
			}
			mazes[i][j] = byte(c)
		}
	}
	for _, p := range mList {
		pointQueue := []point{p}
		for len(pointQueue) != 0 {
			for _, d := range dir {

			}
		}
	}
}