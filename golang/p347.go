package main

import "sort"

type s struct {
	num, count int
}

type list []*s

func (l list) Len() int {
	return len(l)
}

func (l list) Less(i, j int) bool {
	return l[i].count > l[j].count
}

func (l list) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func topKFrequent(nums []int, k int) []int {
	mapSet := make(map[int]int)
	for i := 0; i < len(nums); i ++ {
		mapSet[nums[i]] ++
	}
	l := *(new(list))
	for i, v := range mapSet {
		l = append(l, &s{
			num:   i,
			count: v,
		})
	}
	sort.Sort(l)
	res := []int{}
	for i := 0; i < k; i ++ {
		res = append(res, l[i].num)
	}
	return res
}