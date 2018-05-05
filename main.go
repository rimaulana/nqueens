package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/rimaulana/nqueens/board"
	"github.com/rimaulana/nqueens/solver"
)

var (
	size   int
	sample int
)

// init will always be called before main function
// this will parse the flag passed to the os.Args
func init() {
	flag.IntVar(&size, "size", 4, "is an integer")
	flag.IntVar(&sample, "sample", -1, "is an integer")
}

// main function of the program
func main() {
	// parse the flag passed to os.Args
	flag.Parse()
	// start the clock to time how long does it take
	// for the program to calculate the solution
	solutionStart := time.Now()
	// initialize the solver
	bitwise := solver.New(size)
	// finding solutions
	result := bitwise.Solve()
	// stop the clock
	solutionElapsed := time.Since(solutionStart)

	// initialize board drawer
	boards := board.New(size, sample, &result)
	if sample > -1 {
		var message string
		drawStart := time.Now()
		boardStringStatus := "board"
		if boards.Sample > 1 {
			boardStringStatus += "s are"
		} else {
			boardStringStatus += " is"
		}
		fmt.Print(boards.Draw())
		message = fmt.Sprintf("%d %s drawn in", boards.Sample, boardStringStatus)
		drawElapsed := time.Since(drawStart)
		fmt.Printf("%s %s\n", message, drawElapsed)
	}

	fmt.Printf("There are %d solutions for board size %d\nSolutions are generated in %s\n", len(result), size, solutionElapsed)
}
