// Package board implements board drawing function.
// It defines methods to build the chess board based on the size of the board
package board

import (
	"math/rand"
	"time"
)

// Board defined the stuct of board
type Board struct {
	Size            int
	Data            *[][]int
	Sample          int
	RandomGenerator RandomInterface
}

// RandomInterface provide ability to mock random number generator
type RandomInterface interface {
	Perm(n int) []int
}

// New initialize struct Board
func New(size, sample int, data *[][]int) *Board {
	part := sample
	if sample > len(*data) {
		part = len(*data)
	}
	return &Board{
		Size:            size,
		Data:            data,
		Sample:          part,
		RandomGenerator: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// SetRandom set random number generator, this will allow it
// to be easily mocked
func (brd *Board) SetRandom(rnd RandomInterface) {
	brd.RandomGenerator = rnd
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
func (brd *Board) Draw() string {
	result := ""
	sample := make([]int, 0)
	if brd.Sample != -1 {
		pick := brd.RandomGenerator.Perm(len(*brd.Data))
		sample = pick[:brd.Sample]
	}
	for key := range sample {
		result += header(brd.Size)
		for j := 0; j < len((*brd.Data)[key]); j++ {
			result += row(brd.Size, (*brd.Data)[key][j])
			if j+1 != brd.Size {
				result += separator(brd.Size)
			}
		}
		result += footer(brd.Size)
	}
	return result
}
