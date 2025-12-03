package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"slices"
)

var operations int

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	input := string(f)

	// fmt.Println("task1", getSumOfJoltages(input, 2))
	fmt.Println("task2", getSumOfJoltages(input, 12))
	fmt.Println("operations", operations)
}

func parseInput(input string) []string {
	out := []string{}
	for _, battery := range strings.Split(input, "\n") {
		operations++
		if battery == "" {
			continue
		}
		out = append(out, battery)
	}

	return out
}

func findJoltage(battery string, n int) (joltage string) {
	newBattery := battery
	operations++
	for i := n - 1; i >= 0; i-- {
		fmt.Println(newBattery, i)
		smallerBattery := strings.Split(newBattery[:len(newBattery)-i], "")
		operations++
		// fmt.Println(smallerBattery)
		slices.Sort(smallerBattery)
		operations++
		value := string(smallerBattery[len(smallerBattery)-1])
		operations++
		slicedNewBattery := strings.Split(newBattery, "")
		operations++
		biggestIndex := slices.Index(slicedNewBattery, value)
		operations++
		newBattery = newBattery[biggestIndex+1:]
		operations++
		joltage += value
		operations++
	}
	return joltage
}

func getSumOfJoltages(input string, n int) (sum int) {
	batteries := parseInput(input)
	for _, battery := range batteries {
		joltage := findJoltage(battery, n)
		operations++
		j, _ := strconv.Atoi(joltage)
		operations++
		sum += j
		operations++
	}

	return sum
}
