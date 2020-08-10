package main

var minNum int
var numsArr []int

func findMagicIndex(nums []int) int {
	minNum = -1
	numsArr = nums
	getMagicIndex(0, len(numsArr) - 1)
	return minNum
}

func getMagicIndex(left, right int) {
	for left <= right {
		mid := left + (right - left) / 2
		if mid == numsArr[mid] {
			if minNum == -1 || minNum > mid {
				minNum = mid
			} else {
				if minNum < mid {
					break
				}
			}
			right = mid - 1
		} else {
			getMagicIndex(mid + 1, right)
			right = mid - 1
		}
	}
}