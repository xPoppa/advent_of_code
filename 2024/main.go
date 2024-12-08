package main

import (
	"os"

	"github.com/xPoppa/aoc/2024/day2"
)

func main() {
	home := os.Getenv("HOME")
	day2.Part1(home + "/go/advent_of_code/2024/day2/input.txt")
}
