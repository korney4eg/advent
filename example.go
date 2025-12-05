package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	input := string(f)

	parsedInput := parseInput(input)
	fmt.Println(parsedInput)
}

func parseInput(input string) (field [][]string) {
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		field = append(field, strings.Split(line, ""))

	}
	return field
}
