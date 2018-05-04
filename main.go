package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/rimaulana/nqueens/board"
	"github.com/rimaulana/nqueens/solver"
)

func main() {
	var (
		size    int
		display bool
		sample  int
	)

	flag.IntVar(&size, "size", 4, "is an integer")
	flag.BoolVar(&display, "draw", false, "is a boolean")
	flag.IntVar(&sample, "sample", -1, "is an integer")
	flag.Parse()

	solutionStart := time.Now()
	bitwise := solver.New(size)
	result := bitwise.Solve()
	solutionElapsed := time.Since(solutionStart)

	if display || (sample > -1) {
		var message string
		drawStart := time.Now()
		if sample > -1 {
			fmt.Print(board.Random(size, sample, result))
			message = fmt.Sprintf("%d sample boards were drawn in", sample)
		} else {
			fmt.Print(board.Draw(size, result))
			message = "Boards were drawn in"
		}
		drawElapsed := time.Since(drawStart)
		fmt.Printf("%s %s\n", message, drawElapsed)
	}

	fmt.Printf("There are %d solutions for board size %d\nSolutions are generated in %s\n", len(result), size, solutionElapsed)
}
