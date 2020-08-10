package main

func subarraysDivByK(A []int, K int) int {
	count := 0
	num := 0
	hashSet := make(map[int]int)
	hashSet[0] = 1
	for i := 0; i < len(A); i ++ {
		num += A[i]
		mod := num % K
		if value, ok := hashSet[mod]; ok {
			count += value
			hashSet[mod] ++
		} else {
			hashSet[mod] = 1
		}
	}
	return count
}
