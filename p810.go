package main

func xorGame(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum ^= num
	}
	return sum == 0 && len(nums) % 2 == 0
}