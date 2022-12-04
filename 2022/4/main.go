package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

type ElveRange struct {
	min int
	max int
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	linesList := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		linesList = append(linesList, line)
	}
	count := 0
	for _, line := range linesList {
		firstElve, secondElve := getElvesRanges(line)
		if isOverlapping(firstElve, secondElve) {
			count++
		}

	}

	log.Println(count)
}

func getElvesRanges(line string) (*ElveRange, *ElveRange) {
	firstElve := ElveRange{}
	secondElve := ElveRange{}

	var re = regexp.MustCompile(`\d+`)
	results := re.FindAllStringSubmatch(line, -1)
	firstElve.min, _ = strconv.Atoi(results[0][0])
	firstElve.max, _ = strconv.Atoi(results[1][0])
	secondElve.min, _ = strconv.Atoi(results[2][0])
	secondElve.max, _ = strconv.Atoi(results[3][0])

	return &firstElve, &secondElve
}

func isOverlapping(first, second *ElveRange) bool {
	return (first.min <= second.min && first.max >= second.max) ||
		(first.min >= second.min && first.max <= second.max)
}
