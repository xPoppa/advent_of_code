package day3_test

import (
	"testing"

	"github.com/xPoppa/aoc/2024/day3"
)

func TestPart1(t *testing.T) {
	res := day3.Part1("input_test.txt")

	if res != 161 {
		t.Fatal("Should be 161 got: ", res)
	}
}
