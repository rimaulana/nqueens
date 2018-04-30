// Package board implements board drawing function.
// It defines methods to build the chess board based on the size of the board
package board

import (
	"math/rand"
	"time"
)

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

// header provides function to draw the top edge of the board
func header(size int) string {
	result := "┌"
	for i := 0; i < size; i++ {
		result += "─"
		if i != size-1 {
			result += "┬"
		}
	}
	result += "┐\n"
	return result
}

// footer provides function to draw the bottom edge of the board
func footer(size int) string {
	result := "└"
	for i := 0; i < size; i++ {
		result += "─"
		if i != size-1 {
			result += "┴"
		}
	}
	result += "┘\n"
	return result
}

// separator provides function to draw the connecting line
// between two rows on the board
func separator(size int) string {
	result := "├"
	for i := 0; i < size; i++ {
		result += "─"
		if i != size-1 {
			result += "┼"
		}
	}
	result += "┤\n"
	return result
}

// row provides function to draw row with its queen occupant
func row(size, position int) string {
	result := "│"
	for i := 0; i < size; i++ {
		if power(2, i) == position {
			result += "Q"
		} else {
			result += " "
		}
		result += "│"
	}
	return result + "\n"
}

// Draw help draw all boards from the given map input
func Draw(size int, data map[int][]int) string {
	result := ""
	for key := range data {
		currentPos := 0
		result += header(size)
		data[key] = append(data[key], power(2, size)-1)
		for j := 0; j < len(data[key]); j++ {
			result += row(size, currentPos^data[key][j])
			if j+1 != size {
				result += separator(size)
			}
			currentPos = data[key][j]
		}
		result += footer(size)
	}
	return result
}

// Random help draw radomly selected board(s) from the given map input
func Random(size, sample int, data map[int][]int) string {
	if sample > len(data) {
		return Draw(size, data)
	}
	solutions := make(map[int][]int)
	count := 0
	seed := time.Now()
	rand.Seed(seed.Unix())
	for count < sample {
		pick := rand.Intn(len(data))
		if _, ok := solutions[pick]; !ok {
			solutions[pick] = data[pick]
			count++
		}
	}
	return Draw(size, solutions)

}
