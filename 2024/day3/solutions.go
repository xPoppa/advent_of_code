package day3

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func ParseInput(filename string) []string {
	home := os.Getenv("HOME")
	f, err := os.Open(home + "/go/advent_of_code/2024/day3/" + filename)
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

func Part1(filename string) int {
	input := ParseInput(filename)
	// "mul([0-9],[0-9])" is only valid.
	// Can i fix this with a parser? Or just regex?
	mult := []ToMultiply{}
	for _, str := range input {
		start := strings.Index(str, "mul(")
		if start == -1 {
			continue
		}
		//		reader := strings.NewReader(str)
		//		r, _, _ := reader.ReadRune()
		//		if r == 'm' {
		//			r,_,_ = reader.ReadRune()
		//			if r == 'u' {
		//
		//			}
		//		}
		fmt.Println("The rest of the string with the beginning part", str[start+4:])
		line := str
		for {
			reader := strings.NewReader(line[start+4:])
			toMul, err := readMul(reader)
			if err == io.EOF {
				break
			}
			if err != nil {
				start = strings.Index(line[start+4:], "mul(")
				line = line[start:4]
			}
			mult = append(mult, toMul)

		}

		for _, r := range str[start+4:] {
			unicode.IsDigit(r)
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

	}
	return 0

}

func readMul(r io.RuneReader) (ToMultiply, error) {
	firstNum, _, err := r.ReadRune()
	if err != nil {
		return ToMultiply{}, err
	}
	if !unicode.IsDigit(firstNum) {
		panic("err")
	}
	comma, _, err := r.ReadRune()
	if err != nil {
		return ToMultiply{}, err
	}
	if comma != ',' {
		panic("err")
	}
	sndNum, _, err := r.ReadRune()
	if err != nil {
		return ToMultiply{}, err
	}
	if !unicode.IsDigit(sndNum) {
		panic("err")
	}
	rightParen, _, err := r.ReadRune()
	if err != nil {
		return ToMultiply{}, err
	}
	if rightParen != ')' {
		return ToMultiply{}, errors.New("No right Paren")
	}

	fNum, err := strconv.Atoi(string(firstNum))
	if err != nil {
		return ToMultiply{}, err
	}
	sNum, err := strconv.Atoi(string(sndNum))
	if err != nil {
		return ToMultiply{}, err
	}

	return ToMultiply{firstNum: fNum, secondNum: sNum}, nil

}
