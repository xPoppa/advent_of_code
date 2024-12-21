package day4

import (
	"fmt"

	"github.com/xPoppa/aoc/2024/utils"
)

func Part1(filename string) int {
	input := utils.ReadInput(filename)
	matrix := makeMatrix(input)
	res := countXMAS(input, matrix)

	return res
}

func (m Matrix) String() {
	for k, v := range m {
		fmt.Println("Point: ", k, "With value: ", v)
	}

}

func countXMAS(lines []string, matrix Matrix) (res int) {
	for y, line := range lines {
		for x := range line {
			// Horizontal left to right
			h1 := Point{x, y}
			h2 := Point{x + 1, y}
			h3 := Point{x + 2, y}
			h4 := Point{x + 3, y}
			if matrix[h1] == 'X' && matrix[h2] == 'M' && matrix[h3] == 'A' && matrix[h4] == 'S' {
				res++
			}
			// Reverse horizontal
			hr4 := Point{x, y}
			hr3 := Point{x + 1, y}
			hr2 := Point{x + 2, y}
			hr1 := Point{x + 3, y}
			if matrix[hr1] == 'X' && matrix[hr2] == 'M' && matrix[hr3] == 'A' && matrix[hr4] == 'S' {
				res++
			}
			// Vertical top to bottom
			v1 := Point{x, y}
			v2 := Point{x, y + 1}
			v3 := Point{x, y + 2}
			v4 := Point{x, y + 3}
			if matrix[v1] == 'X' && matrix[v2] == 'M' && matrix[v3] == 'A' && matrix[v4] == 'S' {
				res++
			}
			// Vertical reverse
			vr4 := Point{x, y}
			vr3 := Point{x, y + 1}
			vr2 := Point{x, y + 2}
			vr1 := Point{x, y + 3}
			if matrix[vr1] == 'X' && matrix[vr2] == 'M' && matrix[vr3] == 'A' && matrix[vr4] == 'S' {
				res++
			}
			// diagonal left to right
			d1 := Point{x, y}
			d2 := Point{x + 1, y + 1}
			d3 := Point{x + 2, y + 2}
			d4 := Point{x + 3, y + 3}
			if matrix[d1] == 'X' && matrix[d2] == 'M' && matrix[d3] == 'A' && matrix[d4] == 'S' {
				res++
			}
			// diagonal reverse
			dr4 := Point{x, y}
			dr3 := Point{x + 1, y + 1}
			dr2 := Point{x + 2, y + 2}
			dr1 := Point{x + 3, y + 3}
			if matrix[dr1] == 'X' && matrix[dr2] == 'M' && matrix[dr3] == 'A' && matrix[dr4] == 'S' {
				res++
			}

			// diagonal reverse reverse
			drr1 := Point{x, y}
			drr2 := Point{x + 1, y - 1}
			drr3 := Point{x + 2, y - 2}
			drr4 := Point{x + 3, y - 3}
			if matrix[drr1] == 'X' && matrix[drr2] == 'M' && matrix[drr3] == 'A' && matrix[drr4] == 'S' {
				res++
			}
			// diagonal reverse reverse reverse?
			drrr4 := Point{x, y}
			drrr3 := Point{x + 1, y - 1}
			drrr2 := Point{x + 2, y - 2}
			drrr1 := Point{x + 3, y - 3}
			if matrix[drrr1] == 'X' && matrix[drrr2] == 'M' && matrix[drrr3] == 'A' && matrix[drrr4] == 'S' {
				res++
			}
		}
	}
	return res
}

func makeMatrix(lines []string) Matrix {
	matrix := make(map[Point]rune)
	for y, line := range lines {
		for x, ch := range line {
			matrix[Point{x: x, y: y}] = ch
		}
	}
	return matrix
}

type Point struct {
	x int
	y int
}

type Matrix map[Point]rune
