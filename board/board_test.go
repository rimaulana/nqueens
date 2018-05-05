package board

import "testing"

type mockRand struct {
}

func (mock *mockRand) Perm(n int) []int {
	return make([]int, n)
}

var cases = []struct {
	name   string
	size   int
	sample int
	data   [][]int
	want   string
}{
	{
		name:   "case for size 1",
		size:   1,
		sample: 1,
		data:   [][]int{{1}},
		want:   "┌─┐\n│Q│\n└─┘\n",
	},
	{
		name:   "case for size 4",
		size:   4,
		sample: 1,
		data:   [][]int{{2, 8, 1, 4}, {4, 1, 8, 2}},
		want:   "┌─┬─┬─┬─┐\n│ │Q│ │ │\n├─┼─┼─┼─┤\n│ │ │ │Q│\n├─┼─┼─┼─┤\n│Q│ │ │ │\n├─┼─┼─┼─┤\n│ │ │Q│ │\n└─┴─┴─┴─┘\n",
	},
	{
		name:   "case for size 1 with bigger sample",
		size:   1,
		sample: 2,
		data:   [][]int{{1}},
		want:   "┌─┐\n│Q│\n└─┘\n",
	},
}

func TestBoard_Draw(t *testing.T) {
	for _, tt := range cases {
		brd := New(tt.size, tt.sample, &tt.data)
		brd.SetRandom(&mockRand{})
		t.Run(tt.name, func(t *testing.T) {
			if got := brd.Draw(); got != tt.want {
				t.Errorf("Board.Draw() = %v, want %v", got, tt.want)
			}
		})
	}
}
