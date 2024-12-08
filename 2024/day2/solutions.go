package day2

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/xPoppa/aoc/2024/utils"
)

type Collection struct {
	numbers []int
}

func getData(path string) []Collection {
	lines := utils.ReadInput(path)
	numCollection := []Collection{}
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
	return numCollection
}

func Part1(path string) int {
	numCollection := getData(path)
	counter := 0
	for _, num := range numCollection {
		if num.isSafe() {
			counter++
		}
	}
	fmt.Println("Result is: ", counter)
	return counter
}

func (c Collection) isSafe() bool {
	if c.isDecreasing() {
		return c.onlyDecreasing()
	}
	if c.isIncreasing() {
		return c.onlyIncreasing()
	}
	return false

}

func (c Collection) isDecreasing() bool {
	return c.numbers[0] > c.numbers[1]
}

func (c Collection) isIncreasing() bool {
	return c.numbers[0] < c.numbers[1]
}

func (c Collection) onlyDecreasing() bool {
	for idx := range len(c.numbers) - 1 {
		if c.numbers[idx] < c.numbers[idx+1] {
			return false
		}
		if diff := int(float64(c.numbers[idx] - c.numbers[idx+1])); diff > 3 || diff == 0 {
			return false
		}
	}
	return true
}

func (c Collection) onlyIncreasing() bool {
	for idx := range len(c.numbers) - 1 {
		if c.numbers[idx] > c.numbers[idx+1] {
			return false
		}
		if diff := int((float64(c.numbers[idx] - c.numbers[idx+1]))); diff < -3 || diff == 0 {
			return false
		}
	}
	return true
}
