package solver

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name          string
	totalSolution int
	bit           *Bitwise
	want          [][]int
}{
	{
		name:          "case for size 1",
		totalSolution: 1,
		bit:           New(1),
		want:          [][]int{{1}},
	},
	{
		name:          "case for size 2",
		totalSolution: 0,
		bit:           New(2),
		want:          make([][]int, 0),
	},
	{
		name:          "case for size 4",
		totalSolution: 2,
		bit:           New(4),
		want:          [][]int{{2, 8, 1, 4}, {4, 1, 8, 2}},
	},
	{
		name:          "case for size 5",
		totalSolution: 10,
		bit:           New(5),
		want: [][]int{
			{1, 4, 16, 2, 8},
			{1, 8, 2, 16, 4},
			{2, 8, 1, 4, 16},
			{2, 16, 4, 1, 8},
			{4, 1, 8, 2, 16},
			{4, 16, 2, 8, 1},
			{8, 1, 4, 16, 2},
			{8, 2, 16, 4, 1},
			{16, 2, 8, 1, 4},
			{16, 4, 1, 8, 2},
		},
	},
}

func TestBitwise_Solve(t *testing.T) {
	for _, fixture := range cases {
		t.Run(fixture.name, func(t *testing.T) {
			got := fixture.bit.Solve()
			if !reflect.DeepEqual(got, fixture.want) {
				t.Errorf("Bitwise.Solve() = %v, want %v", got, fixture.want)
			}
			if len(got) != fixture.totalSolution {
				t.Errorf("Total of solutions = %d, want %d", len(got), fixture.totalSolution)
			}
		})
	}
}
