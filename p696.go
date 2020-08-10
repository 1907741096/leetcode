package main

func countBinarySubstrings(s string) int {
	num0, num1 := 0, 0
	stack0, stack1 := []int{}, []int{}
	for i, v := range s {
		if v == '0' {
			if len(stack1) != 0 {
				num0 ++
				if len(stack1) == 1 {
					stack1 = stack1[:len(stack1) - 1]
				} else {
					if stack1[len(stack1) - 1] == stack1[len(stack1) - 2] + 1 {
						stack1 = stack1[:len(stack1) - 1]
					} else {
						stack1 = []int{}
					}
				}
			}
			stack0 = append(stack0, i)
		} else {
			if len(stack0) != 0 {
				num1 ++
				if len(stack0) == 1 {
					stack0 = stack0[:len(stack0) - 1]
				} else {
					if stack0[len(stack0) - 1] == stack0[len(stack0) - 2] + 1 {
						stack0 = stack0[:len(stack0) - 1]
					} else {
						stack0 = []int{}
					}
				}
			}
			stack1 = append(stack1, i)
		}
	}
	return num0 + num1
}