package main

import (
	"fmt"
	"os"
	"strings"
)

type position struct {
	x int
	y int
}

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	input := string(f)

	field := parseInput(input)
	fmt.Println("task1", len(findAvailableRolls(field)))
	fmt.Println("task2", cleanupField(field))
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

func countNeighbors(field [][]string, pos position) int {
	neighboors := 0
	for y := pos.y - 1; y <= pos.y+1; y++ {
		if y < 0 || y >= len(field) {
			continue
		}
		for x := pos.x - 1; x <= pos.x+1; x++ {
			if x < 0 || x >= len(field[y]) || (x == pos.x && y == pos.y) {
				continue
			}
			if field[y][x] == "@" {
				neighboors++
			}
		}
	}
	return neighboors
}

func findAvailableRolls(field [][]string) (points []position) {
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			currentField := field[y][x]
			neighboors := countNeighbors(field, position{x: x, y: y})
			if currentField == "@" && neighboors < 4 {
				points = append(points, position{x: x, y: y})
			}
		}
	}
	return points
}

func cleanupField(field [][]string) (cleans int) {
	for {
		pointsToClean := findAvailableRolls(field)
		cleans += len(pointsToClean)
		if len(pointsToClean) == 0 {
			break
		}
		for _, point := range pointsToClean {
			field[point.y][point.x] = "x"
		}
	}
	return cleans
}
