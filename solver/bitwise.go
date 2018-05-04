// Package solver implements n-queens solver using bitwise approach and recursion.
// It defines methods to find the solution for n-queens problem
package solver

// Bitwise define the struct
type Bitwise struct {
	size int
}

// New initiate new instance of Bitwise
func New(size int) *Bitwise {
	return &Bitwise{size: size}
}

// power provides function to find x power of n which return integer.
// this function is utilized as a replacement of Math.Pow which
// unlike Math.Pow, this function accept Int and return Int as a product
func power(n, x int) int {
	result := 1
	for i := 0; i < x; i++ {
		result *= n
	}
	return result
}

// isDone check whether there is no more safe tile to add new queens
// when this condition is reached, it means that a solution has been
// found and needs to be added to the collection of solutions
func isDone(done int, col int, count *int, stack []int, result map[int][]int) bool {
	// if done == col means that there is no more safe tile to add
	// new queen
	if done == col {
		// by default map is passed by reference so we can easily change its
		// content by adding new solution
		result[*count] = stack
		// increase the number of solution find. the variable is passed to
		// function as a pointer so that other function will see the change too
		*count++
		return true
	}
	return false
}

// Solve finds the solution of the n-queens problem with the input of the
// board size. It will find all the possible solution of the problem in
// a map[int][]int datatype.
func (bit *Bitwise) Solve() (result map[int][]int) {
	// initialize count of solution and map to accommodate the solutions
	resultCount := 0
	result = make(map[int][]int)

	// set condition where there is no more safe tile for new queen
	finished := power(2, bit.size) - 1

	// defined the closure so that it can be called inside itself
	// as a recursion function
	var walk func(param1, param2, param3 int, param4 []int)

	// define the logic of walk function
	walk = func(leftDiagonal, column, rightDiagonal int, stack []int) {
		if isDone(finished, column, &resultCount, stack, result) {
			return
		}
		if column != 0 {
			stack = append(stack, column)
		}
		position := ^(leftDiagonal | column | rightDiagonal)
		for (position & finished) != 0 {
			candidate := position & -position
			position -= candidate
			walk((leftDiagonal|candidate)>>1, column|candidate, (rightDiagonal|candidate)<<1, stack)
		}
	}
	walk(0, 0, 0, make([]int, 0))
	return result
}
