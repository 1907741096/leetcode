package main

func validateStackSequences(pushed []int, popped []int) bool {
	i, j := 0, 0
	stack := []int{}
	for i < len(pushed) || j < len(popped) {
		if len(stack) == 0 {
			stack = append(stack, pushed[i])
			i ++
		}
		if i == len(pushed) && stack[len(stack) - 1] != popped[j] {
			return false
		}
		if stack[len(stack) - 1] == popped[j] {
			stack = stack[0: len(stack) - 1]
			j ++
		} else {
			stack = append(stack, pushed[i])
			i ++
		}
	}
	return true
}