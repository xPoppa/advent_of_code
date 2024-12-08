package day2

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/xPoppa/aoc/2024/utils"
)

type Collection struct {
	numbers []int
}

func Part1() {
	lines := utils.ReadInput("/home/romano/go/advent_of_code/2024/day2/input.txt")
	numCollection := []Collection{}
	counter := 0
	for _, line := range lines {
		splittedString := strings.Split(line, " ")
		nums := make([]int, len(splittedString))
		for idx, sNum := range splittedString {
			num, err := strconv.Atoi(sNum)
			if err != nil {
				log.Fatal("Cannot turn string to number: \n", sNum)
			}
			nums[idx] = num
		}
		numCollection = append(numCollection, Collection{numbers: nums})
	}

	for _, num := range numCollection {
		if num.isSafe() {
			counter++
		}
	}

	fmt.Println("Result is: ", counter)
}

func (c Collection) isSafe() bool {
	return c.onlyDecreasing() || c.onlyIncreasing()

}

func (c Collection) onlyDecreasing() bool {
	for idx := range len(c.numbers) - 2 {
		if c.numbers[idx] > c.numbers[idx+1] {
			return false
		}
		if int(math.Abs(float64(c.numbers[idx]-c.numbers[idx+1]))) > 2 {
			return false
		}
	}
	return true
}

func (c Collection) onlyIncreasing() bool {
	for idx := range len(c.numbers) - 2 {
		if c.numbers[idx] < c.numbers[idx+1] {
			return false
		}
		if int(math.Abs(float64(c.numbers[idx]-c.numbers[idx+1]))) > 3 {
			return false
		}
	}
	return true
}
