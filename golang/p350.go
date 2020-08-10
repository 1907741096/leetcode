package main

func intersect(nums1 []int, nums2 []int) []int {
	hashMap := make(map[int]int)
	for _, v := range nums1 {
		if _, ok := hashMap[v]; ok {
			hashMap[v] ++
		} else {
			hashMap[v] = 1
		}
	}

	resArr := []int{}
	for _, v := range nums2 {
		if n, ok := hashMap[v]; !ok || n == 0 {
			continue
		} else {
			hashMap[v] --
			resArr = append(resArr, v)
		}
	}

	return resArr
}