package board

import (
	"testing"
)

func Test_power(t *testing.T) {
	type args struct {
		n int
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"power_of_two",
			args{n: 2, x: 2},
			4,
		},
		{
			"power_of_three",
			args{n: 3, x: 2},
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := power(tt.args.n, tt.args.x); got != tt.want {
				t.Errorf("power() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_header(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"size 1",
			args{size: 1},
			"┌─┐\n",
		},
		{
			"size 2",
			args{size: 4},
			"┌─┬─┬─┬─┐\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := header(tt.args.size); got != tt.want {
				t.Errorf("header() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_footer(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"size 1",
			args{size: 1},
			"└─┘\n",
		},
		{
			"size 2",
			args{size: 4},
			"└─┴─┴─┴─┘\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := footer(tt.args.size); got != tt.want {
				t.Errorf("footer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_separator(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"size 1",
			args{size: 8},
			"├─┼─┼─┼─┼─┼─┼─┼─┤\n",
		},
		{
			"size 2",
			args{size: 12},
			"├─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤\n",
		},
		{
			"size 3",
			args{size: 2},
			"├─┼─┤\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := separator(tt.args.size); got != tt.want {
				t.Errorf("separator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_row(t *testing.T) {
	type args struct {
		size     int
		position int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"size 1 pos 1",
			args{size: 1, position: 1},
			"│Q│\n",
		},
		{
			"size 5 pos 5",
			args{size: 5, position: 16},
			"│ │ │ │ │Q│\n",
		},
		{
			"size 4 pos 2",
			args{size: 4, position: 2},
			"│ │Q│ │ │\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := row(tt.args.size, tt.args.position); got != tt.want {
				t.Errorf("row() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_draw(t *testing.T) {
	type args struct {
		size int
		data map[int][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"3x3",
			args{
				size: 3,
				data: map[int][]int{
					0: []int{1, 3},
					1: []int{4, 6},
				},
			},
			"┌─┬─┬─┐\n│Q│ │ │\n├─┼─┼─┤\n│ │Q│ │\n├─┼─┼─┤\n│ │ │Q│\n└─┴─┴─┘\n┌─┬─┬─┐\n│ │ │Q│\n├─┼─┼─┤\n│ │Q│ │\n├─┼─┼─┤\n│Q│ │ │\n└─┴─┴─┘\n",
		},
		{
			"1x1",
			args{
				size: 1,
				data: map[int][]int{
					0: []int{},
				},
			},
			"┌─┐\n│Q│\n└─┘\n",
		},
		{
			"2x2",
			args{
				size: 2,
				data: map[int][]int{
					0: []int{1},
					1: []int{2},
				},
			},
			"┌─┬─┐\n│Q│ │\n├─┼─┤\n│ │Q│\n└─┴─┘\n┌─┬─┐\n│ │Q│\n├─┼─┤\n│Q│ │\n└─┴─┘\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Draw(tt.args.size, tt.args.data); got != tt.want {
				t.Errorf("draw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_random(t *testing.T) {
	type args struct {
		size   int
		sample int
		data   map[int][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1x1",
			args{
				size:   1,
				sample: 1,
				data: map[int][]int{
					0: []int{},
				},
			},
			"┌─┐\n│Q│\n└─┘\n",
		},
		{
			"2x2",
			args{
				size:   2,
				sample: 1,
				data: map[int][]int{
					0: []int{1},
					1: []int{1},
				},
			},
			"┌─┬─┐\n│Q│ │\n├─┼─┤\n│ │Q│\n└─┴─┘\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Random(tt.args.size, tt.args.sample, tt.args.data); got != tt.want {
				t.Errorf("random() = %v, want %v", got, tt.want)
			}
		})
	}
}
