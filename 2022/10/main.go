package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var register = 1
var cyclesCount = 0

var monitor = []string{
	"........................................",
	"........................................",
	"........................................",
	"........................................",
	"........................................",
	"........................................",
}

type CommandResult struct {
	cyclesCount   int
	registerDelta int
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		result := getResult(&line)

		for i := 0; i < result.cyclesCount; i++ {
			lineIndex := cyclesCount / 40
			pointIndex := cyclesCount % 40
			cyclesCount++

			if lineIndex > len(monitor)-1 {
				break
			}

			monitorLineStr := monitor[lineIndex]

			monitorLineArr := strings.Split(monitorLineStr, "")

			if register-1 <= pointIndex && register+1 >= pointIndex {
				monitorLineArr[pointIndex] = "#"
			}

			monitor[lineIndex] = strings.Join(monitorLineArr, "")

			if i == 1 {
				register += result.registerDelta
			}
		}

	}

	draw(&monitor)
	//fmt.Printf("%v\n", monitor)
}

func getResult(line *string) *CommandResult {
	if *line == "noop" {
		return &CommandResult{
			cyclesCount:   1,
			registerDelta: 0,
		}
	}

	registerDelta, _ := strconv.Atoi(strings.Replace(*line, "addx ", "", -1))

	return &CommandResult{
		cyclesCount:   2,
		registerDelta: registerDelta,
	}
}

func getIsSuit(index int) bool {
	if index < 0 {
		return false
	}

	return index <= 39
}

func draw(data *[]string) {
	for _, line := range *data {
		fmt.Println(line)
	}
}
