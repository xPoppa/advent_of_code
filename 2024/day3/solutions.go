package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput() []string {
	home := os.Getenv("HOME")
	f, err := os.Open(home + "/go/advent_of_code/2024/day3/input.txt")
	if err != nil {
		log.Fatal("Errored out while reading file: ", err)
	}
	rawInput := bufio.NewScanner(f)

	input := []string{}
	for rawInput.Scan() {
		line := rawInput.Text()
		input = append(input, line)
	}

	return input
}

//xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))

type ToMultiply struct {
	firstNum  int
	secondNum int
}

func Part1() {
	input := parseInput()
	// "mul([0-9],[0-9])" is only valid.
	// Can i fix this with a parser? Or just regex?
	mult := []ToMultiply{}
	for _, str := range input {
		start := strings.Index(str, "mul(")
		if start == -1 {
			continue
		}
		firstNumStr := rune(str[start+1])
		fmt.Println("The first number: ", firstNumStr)
		comma := rune(str[start+2])
		secondNumStr := rune(str[start+3])
		rightParen := rune(str[start+4])
		firstNum, err := strconv.Atoi(string(firstNumStr))
		secondNum, err := strconv.Atoi(string(secondNumStr))
		if comma == ',' && err == nil && rightParen == ')' {
			mult = append(mult, ToMultiply{firstNum: firstNum, secondNum: secondNum})
		}

		newStr := str[start+5:]
	}

}
