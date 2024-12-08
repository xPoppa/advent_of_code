package day1

import (
	"bufio"
	"fmt"
	"math"
	"sort"

	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("solution part2 is: ", part2())
}

func part2() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("The file is not apparent", err)
	}
	defer file.Close()

	lines := []string{}

	rd := bufio.NewReader(file)
	first_row := []string{}
	last_row := []string{}
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("read file line error: %v", err)
			return 0
		}
		lines = append(lines, line)
	}

	for _, line := range lines {
		splittedString := strings.Split(line, " ")
		first_row = append(first_row, splittedString[0])
		last_row = append(last_row, strings.Trim(splittedString[3], "\n"))
	}

	firstIntRow := []int{}
	lastIntRow := []int{}

	for _, s := range first_row {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal("Dont deal with it")
		}
		firstIntRow = append(firstIntRow, num)
	}
	for _, s := range last_row {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal("Dont deal with it: ", err)
		}
		lastIntRow = append(lastIntRow, num)
	}

	return totalPart2(CountMembersAndAmounts(lastIntRow, firstIntRow))
}

type NumberAndAmount struct {
	number       int
	memberAmount int
}

func totalPart2(nas []NumberAndAmount) int {
	total := 0
	for _, na := range nas {
		total += na.number * na.memberAmount
	}
	return total
}

func CountMembersAndAmounts(right []int, left []int) []NumberAndAmount {
	na := []NumberAndAmount{}
	for _, leftEl := range left {
		na = append(na, NumberAndAmount{number: leftEl, memberAmount: isMemberAmount(right, leftEl)})
	}
	return na
}

func isMemberAmount(s []int, maybeMem int) int {
	members := 0
	for _, el := range s {
		if el == maybeMem {
			members++
		}
	}
	return members
}

func part1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("The file is not apparent", err)
	}
	defer file.Close()

	lines := []string{}

	rd := bufio.NewReader(file)
	first_row := []string{}
	last_row := []string{}
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("read file line error: %v", err)
			return 0
		}
		lines = append(lines, line)
	}

	for _, line := range lines {
		splittedString := strings.Split(line, " ")
		first_row = append(first_row, splittedString[0])
		last_row = append(last_row, strings.Trim(splittedString[3], "\n"))
	}

	firstIntRow := []int{}
	lastIntRow := []int{}

	for _, s := range first_row {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal("Dont deal with it")
		}
		firstIntRow = append(firstIntRow, num)
	}
	for _, s := range last_row {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal("Dont deal with it: ", err)
		}
		lastIntRow = append(lastIntRow, num)
	}
	sort.Ints(firstIntRow)
	sort.Ints(lastIntRow)

	score := []int{}

	for idx, left := range firstIntRow {
		if len(lastIntRow)-1 < idx {
			break
		}
		right := lastIntRow[idx]
		fmt.Println("Add to score: ", right-left)
		fmt.Println("Left value: ", left)
		fmt.Println("Right value: ", right)
		score = append(score, int(math.Abs(float64(right-left))))
		continue
	}

	total := 0

	for _, s := range score {
		total += s
	}

	fmt.Println("The total is: ", total)
	return total
}
