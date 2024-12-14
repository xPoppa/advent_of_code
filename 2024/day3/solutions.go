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
		before, after, found := strings.Cut(str, "mul(")
		fmt.Println("The after is: ", after)
		fmt.Println("Before: ", before)
		if !found {
			fmt.Println("Miauw")
			continue
		}
		line := after
		fmt.Println("The full line: \n\t", str)
		for {
			fmt.Println("The line I am checking: \n\t", line)
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
				fmt.Println("ERROR After assigning new line: ", line)
				continue
			}
			mult = append(mult, toMul)
			_, line, found = strings.Cut(line, "mul(")
			fmt.Println("After checking assigning new line inside NORMAL")
		}
	}
	res := 0
	for _, mul := range mult {
		res += mul.Product()
	}
	return res
}

func (m ToMultiply) Product() int {
	fmt.Printf("%d * %d\n", m.firstNum, m.secondNum)
	return m.firstNum * m.secondNum
}

func readMul(r *bufio.Reader) (ToMultiply, error) {
	firstNum, _, err := r.ReadRune()
	if err != nil {
		return ToMultiply{}, err
	}
	if !unicode.IsDigit(firstNum) {
		return ToMultiply{}, errors.New("Not the right character")
	}
	comma, _, err := r.ReadRune()
	if err != nil {
		return ToMultiply{}, err
	}
	if comma != ',' {
		return ToMultiply{}, errors.New("Not the right character")
	}
	sndNum, _, err := r.ReadRune()
	if err != nil {
		return ToMultiply{}, err
	}
	if !unicode.IsDigit(sndNum) {
		return ToMultiply{}, errors.New("Not the right character")
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
