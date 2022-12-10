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

	//var commandList []string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		result := getResult(&line)
		//fmt.Printf("%v\n", result)

		//cyclesCount += result.cyclesCount

		//fmt.Printf("%d\n", cyclesCount)

		for i := 0; i < result.cyclesCount; i++ {
			//fmt.Printf("%s\n", line)
			cyclesCount++

			if cyclesCount == 20 ||
				cyclesCount == 60 ||
				cyclesCount == 100 ||
				cyclesCount == 140 ||
				cyclesCount == 180 ||
				cyclesCount == 220 {

				sum += register * cyclesCount
				//fmt.Printf("---> %d, %d\n", register*cyclesCount, register)
			} else {
				//fmt.Printf("%d, %d\n", cyclesCount, register)
			}

			if i == 1 {
				register += result.registerDelta
			}

		}

		//commandList = append(commandList, line)
	}

	fmt.Printf(">>>> %d\n", sum)
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
