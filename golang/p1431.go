package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	arr := make([]bool, len(candies))
	for _, v := range candies {
		if max < v {
			max = v
		}
	}
	for i, v := range candies {
		if v + extraCandies < max {
			arr[i] = false
		} else {
			arr[i] = true
		}
	}
	return arr
}