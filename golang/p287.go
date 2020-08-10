package main

func findDuplicate(nums []int) int {
	count := len(nums)
	slow := 0
	fast := 1
	for {
		if slow == fast {
			fast ++
			continue
		}
		if slow >= count {
			slow -= count
		}
		if fast >= count {
			fast -= count
		}
		if nums[slow] == nums[fast] {
			return nums[fast]
		}
		slow ++
		fast += 2
	}
}
