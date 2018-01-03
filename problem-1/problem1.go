package problem1

func solution(stack []int) []int {
	newStack := make([]int, 0)

	indexMinFirstHalf := 0
	indexMaxSecondHalf := len(stack) - 1

	for len(newStack) < len(stack) {
		if len(stack)-len(newStack) == 1 {
			newStack = append(newStack, stack[indexMaxSecondHalf])
			break
		}

		newStack = append(newStack, stack[indexMinFirstHalf])
		indexMinFirstHalf++

		newStack = append(newStack, stack[indexMaxSecondHalf])
		indexMaxSecondHalf--
	}

	return newStack
}
