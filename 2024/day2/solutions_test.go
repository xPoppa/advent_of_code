package day2_test

import (
	"os"
	"testing"

	"github.com/xPoppa/aoc/2024/day2"
)

func TestIsSafe(t *testing.T) {
	home := os.Getenv("HOME")
	input := home + "/go/advent_of_code/2024/day2/test.txt"
	res := day2.Part1(input)

	if res != 2 {
		t.Fatal("Didn't get the right amount should be 2 but is: ", res)
	}
}

func TestIsPDSafe(t *testing.T) {
	home := os.Getenv("HOME")
	input := home + "/go/advent_of_code/2024/day2/test.txt"
	res := day2.Part2(input)

	if res != 4 {
		t.Fatal("Didn't get the right amount should be 4 but is: ", res)
	}
}
