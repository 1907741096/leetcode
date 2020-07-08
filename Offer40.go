package main

func getLeastNumbers(arr []int, k int) []int {
	sortArr := myQuickSort(arr)
	return sortArr[:k]
}

func myQuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	leftArr := []int{}
	rightArr := []int{}
	for i := 1; i < len(arr); i ++ {
		if arr[i] <= arr[0] {
			leftArr = append(leftArr, arr[i])
		} else {
			rightArr = append(rightArr, arr[i])
		}
	}
	return append(append(myQuickSort(leftArr), arr[0]), myQuickSort(rightArr)...)
}