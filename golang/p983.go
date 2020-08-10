package main

func mincostTickets(days []int, costs []int) int {
	arrs := make(map[int]map[string]int)
	for i, day := range days {
		if i == 0 {
			arrs[day] = make(map[string]int)
			arrs[day]["one"] = 1
			arrs[day]["seven"] = 0
			arrs[day]["thirty"] = 0
			arrs[day]["money"] = arrs[day]["one"] * costs[0] + arrs[day]["seven"] * costs[1] + arrs[day]["thirty"] * costs[2]
		} else {
			arrs[day] = make(map[string]int)
			arrs[day]["one"] = 1
			arrs[day]["seven"] = 0
			arrs[day]["thirty"] = 0
			arrs[day]["money"] = arrs[day]["one"] * costs[0] + arrs[day]["seven"] * costs[1] + arrs[day]["thirty"] * costs[2]
		}
	}
	//for day, arr := range arrs {
	//
	//}
	return arrs[days[len(arrs)]]["money"]
}