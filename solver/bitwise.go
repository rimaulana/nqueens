package solver

func power(n, x int) int {
	result := 1
	for i := 0; i < x; i++ {
		result *= n
	}
	return result
}

func Bitwise(size int) (result map[int][]int) {
	resultCount := 0
	result = make(map[int][]int)

	finished := power(2, size) - 1
	var recurse func(param1, param2, param3 int, param4 []int)
	recurse = func(leftDiagonal, column, rightDiagonal int, stack []int) {
		if finished == column {
			result[resultCount] = stack
			resultCount++
			return
		}
		if column != 0 {
			stack = append(stack, column)
		}
		position := ^(leftDiagonal | column | rightDiagonal)
		for (position & finished) != 0 {
			candidate := position & -position
			position -= candidate
			recurse((leftDiagonal|candidate)>>1, column|candidate, (rightDiagonal|candidate)<<1, stack)
		}
	}
	recurse(0, 0, 0, make([]int, 0))
	return result
}
