package main

import "strconv"

func threeSum(nums []int) [][]int {
	arr := [][]int{}
	arrSet := make(map[string]int)
	hastSet := make(map[int]int)
	for i, v := range nums {
		hastSet[v] = i
	}
	length := len(nums)
	for i := 0; i < length; i ++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		for j := i + 1; j < length; j ++ {
			k := 0 - nums[i] - nums[j]
			if index, ok := hastSet[k]; ok && index != i && index != j {
				str := getThreeSortString(nums[i], nums[j], k)
				if _, ok := arrSet[str]; !ok {
					arrSet[str] = 1
					arr = append(arr, []int{nums[i], nums[j], k})
				}
			}
		}
	}
	return arr
}

func getThreeSortString(i, j, k int) string {
	var str string
	if i > j {
		if k > i {
			str = strconv.Itoa(k) + strconv.Itoa(i) + strconv.Itoa(j)
		} else {
			if k > j {
				str = strconv.Itoa(i) + strconv.Itoa(k) + strconv.Itoa(j)
			} else {
				str = strconv.Itoa(i) + strconv.Itoa(j) + strconv.Itoa(k)
			}
		}
	} else {
		if k > j {
			str = strconv.Itoa(k) + strconv.Itoa(j) + strconv.Itoa(i)
		} else {
			if k > i {
				str = strconv.Itoa(j) + strconv.Itoa(k) + strconv.Itoa(i)
			} else {
				str = strconv.Itoa(j) + strconv.Itoa(i) + strconv.Itoa(k)
			}
		}
	}
	return str
}
