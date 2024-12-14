package day3

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var NOT_RIGHT_CHAR error = errors.New("Not the right character")

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
	firstNum   int
	secondNum  int
	isDisabled bool
}

func Part1(filename string) int {
	input := ParseInput(filename)
	// "mul([0-9],[0-9])" is only valid.
	// Can i fix this with a parser? Or just regex?
	mult := []ToMultiply{}
	for _, str := range input {
		_, after, found := strings.Cut(str, "mul(")
		if !found {
			continue
		}
		line := after
		for {
			reader := bufio.NewReader(strings.NewReader(line))
			toMul, err := readMul(reader)
			if err == io.EOF {
				break
			}
			if err != nil {
				_, line, found = strings.Cut(line, "mul(")
				if !found {
					break
				}
				continue
			}

			mult = append(mult, toMul)
			_, line, found = strings.Cut(line, "mul(")
		}
	}
	res := 0
	for _, mul := range mult {
		res += mul.Product()
	}
	return res
}

func Part2(filename string) int {
	input := ParseInput(filename)
	// "mul([0-9],[0-9])" is only valid.
	// Can i fix this with a parser? Or just regex?
	mult := []ToMultiply{}
	for _, str := range input {
		before, after, found := strings.Cut(str, "mul(")
		if !found {
			continue
		}
		_, _, disable := strings.Cut(before, "don't()")
		line := after
		for {
			reader := bufio.NewReader(strings.NewReader(line))
			toMul, err := readMul(reader)
			toMul.isDisabled = disable
			if err == io.EOF {
				break
			}
			if err != nil {
				before, line, found = strings.Cut(line, "mul(")
				if !found {
					break
				}
				_, _, disable = strings.Cut(line, "don't()")
				continue
			}

			mult = append(mult, toMul)
			_, line, found = strings.Cut(line, "mul(")
			_, _, disable = strings.Cut(line, "don't()")
		}
	}
	res := 0
	for _, mul := range mult {
		res += mul.Product()
	}
	return res

}

func (m ToMultiply) Product() int {
	if m.isDisabled {
		return 1
	}
	return m.firstNum * m.secondNum
}

func readMul(r *bufio.Reader) (ToMultiply, error) {
	firstNum, err := readNumber(r)
	if err != nil {
		return ToMultiply{}, err
	}

	comma, _, err := r.ReadRune()
	if err != nil {
		return ToMultiply{}, err
	}
	if comma != ',' {
		return ToMultiply{}, errors.New("Not the right character")
	}

	sndNum, err := readNumber(r)
	if err != nil {
		return ToMultiply{}, err
	}

	rightParen, _, err := r.ReadRune()
	if err != nil {
		return ToMultiply{}, err
	}
	if rightParen != ')' {
		return ToMultiply{}, errors.New("No right Paren")
	}

	fNum, err := strconv.Atoi(firstNum)
	if err != nil {
		return ToMultiply{}, err
	}
	sNum, err := strconv.Atoi(sndNum)
	if err != nil {
		return ToMultiply{}, err
	}
	return ToMultiply{firstNum: fNum, secondNum: sNum}, nil
}

func readNumber(r *bufio.Reader) (string, error) {
	firstNum, _, err := r.ReadRune()
	if err != nil {
		return "", err
	}
	if !unicode.IsDigit(firstNum) {
		return "", NOT_RIGHT_CHAR
	}
	nums := string(firstNum)

	for {
		maybeNum, _, err := r.ReadRune()
		if err != nil {
			r.UnreadRune()
			break
		}
		if !unicode.IsDigit(maybeNum) {
			r.UnreadRune()
			break
		}
		nums = nums + string(maybeNum)
	}

	return nums, nil
}
