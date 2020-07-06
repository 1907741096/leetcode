package main

func longestValidParentheses(s string) int {
	stack := []int{-1}
	max := 0
	for i, v := range s {
		if v == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[0: len(stack) - 1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				now := i - stack[len(stack)-1]
				if now > max {
					max = now
				}
			}
		}
	}
	return max
}