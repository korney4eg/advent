package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	input := string(f)

	numbers, operations := parseInput(input)
	fmt.Println("task1", makeOperations(numbers, operations))
	fmt.Println("task2", makeLeftToRightOperations(numbers, operations))
}

func parseInput(input string) (numbers [][]string, operations []string) {

	splitedLine := strings.Split(strings.TrimRight(input, "\n"), "\n")
	spaces := []int{}
	for i := 0; i < len(splitedLine[0]); i++ {
		spaced := true
		for j := 0; j < len(splitedLine); j++ {
			if string(splitedLine[j][i]) != " " {
				spaced = false
			}
		}
		if spaced {
			spaces = append(spaces, i)
		}
	}
	numbers = make([][]string, len(spaces)+1)
	for _, line := range splitedLine[0 : len(splitedLine)-1] {
		prevSpace := 0
		for i, space := range spaces {
			numbers[i] = append(numbers[i], line[prevSpace:space])
			prevSpace = space
		}
		numbers[len(numbers)-1] = append(numbers[len(numbers)-1], line[prevSpace:])

	}

	for _, word := range strings.Split(splitedLine[len(splitedLine)-1], " ") {
		if word == "" {
			continue
		}
		operations = append(operations, word)
	}
	return numbers, operations
}

func makeOperations(numbers [][]string, operations []string) int {
	total := 0
	for i, problems := range numbers {
		subtotal := 0
		if operations[i] == "*" {
			subtotal = 1
		}
		for _, nums := range problems {
			num, _ := strconv.Atoi(strings.TrimSpace(nums))
			if operations[i] == "*" {
				subtotal *= num
			} else {
				subtotal += num
			}
		}
		total += subtotal
	}
	return total
}

func makeLeftToRightOperations(numbers [][]string, operations []string) int {
	total := 0
	for i, problem := range numbers {
		subtotal := 0
		if operations[i] == "*" {
			subtotal = 1
		}
		allNums := []string{}
		for x := 0; x < len(problem[0]); x++ {
			nums := ""
			for _, p := range problem {
				nums += string(p[x])
			}
			if strings.TrimSpace(nums) == "" {
				continue
			}
			allNums = append(allNums, nums)
		}
		for _, nums := range allNums {
			num, _ := strconv.Atoi(strings.TrimSpace(nums))
			if operations[i] == "*" {
				subtotal *= num
			} else {
				subtotal += num
			}
		}
		total += subtotal
	}
	return total
}
