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

// reformat will format the output from Solve, all the numbers will
// be performed XOR operation with previous number to produce power of 2 number
// which in its binary representation, the 1 bit
// represent the location of the queens
func reformat(data []int) (result []int) {
	currentPosition := 0
	data = append(data, power(2, len(data)+1)-1)
	for i := 0; i < len(data); i++ {
		result = append(result, currentPosition^data[i])
		currentPosition = data[i]
	}
	return result
}

// Solve finds the solution of the n-queens problem with the input of the
// board size. It will find all the possible solution of the problem in
// a map[int][]int datatype.
func (bit *Bitwise) Solve() [][]int {
	// initialize count of solution and map to accommodate the solutions
	result := make([][]int, 0)

	// set condition where there is no more safe tile for new queen
	finished := power(2, bit.size) - 1

	// defined the closure so that it can be called inside itself
	// as a recursion function
	var walk func(param1, param2, param3 int, param4 []int)

	// define the logic of walk function
	walk = func(leftDiagonal, column, rightDiagonal int, stack []int) {
		// when all the column has been occupied by queens, it means
		// that new solution has been found.
		if finished == column {
			// add new solution to the collection of solutions
			result = append(result, reformat(stack))
			return
		}
		// if it is not the initial walk call, then add the new occupied
		// tile into stack
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
	// call the first walk on the tile
	walk(0, 0, 0, make([]int, 0))
	return result
}
