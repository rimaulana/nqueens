package board

import (
	"math/rand"
	"time"
)

func power(n, x int) int {
	result := 1
	for i := 0; i < x; i++ {
		result *= n
	}
	return result
}

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
