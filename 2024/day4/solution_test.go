package day4_test

import (
	"testing"

	"github.com/xPoppa/aoc/2024/day4"
)

func TestPart1(t *testing.T) {
	res := day4.Part1("input_test.txt")

	if res != 18 {
		t.Fatal("Ow my god not the right amount, you had: ", res)
	}

}
