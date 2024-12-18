package day4

import (
	"fmt"

	"github.com/xPoppa/aoc/2024/utils"
)

func Part1(filename string) int {
	input := utils.ReadInput(filename)

	return 0
}

func countXMAS(lines []string) (res int) {
	counter := 0
	for _, line := range lines {
		counter = counter + findXMAS(line)
	}
	return counter
}

func findXMAS(line string) (res int) {
	counter := 0
	for i := range line {
		if i+3 > len(line) {
			break
		}
		if line[i] == 'X' && line[i+1] == 'M' && line[i+2] == 'A' && line[i+3] == 'S' {
			counter++
		}
		if line[i+3] == 'X' && line[i+2] == 'M' && line[i+1] == 'A' && line[i] == 'S' {
			counter++
		}
	}
	return counter

}

func makeHorizontal(lines []string) (horizontal []string) {
	for _, line := range lines {

	}
}
