package main

import (
	"log"
	"os"
	"strings"
)

type position struct {
	x, y int
}

var robot = position{}

func parseInput(input string) (field [][]string, instructions []string) {
	isInstructions := false
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		if line == "" {
			isInstructions = true
			continue
		}
		if isInstructions {
			instructions = append(instructions, strings.Split(line, "")...)
		} else {
			field = append(field, strings.Split(line, ""))
			x := strings.Index(line, "@")
			if x != -1 {
				robot.x = x
				robot.y = y
			}
		}
	}
	return field, instructions

}
func parseInputWide(input string) (field [][]string, instructions []string) {
	isInstructions := false
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		if line == "" {
			isInstructions = true
			continue
		}
		if isInstructions {
			instructions = append(instructions, strings.Split(line, "")...)
		} else {
			l := []string{}
			for _, cell := range line {
				stringedCell := string(cell)
				if stringedCell == "O" {
					l = append(l, "[", "]")
				} else if stringedCell == "@" {
					robot.x = len(l)
					robot.y = y
					l = append(l, "@", ".")
				} else {
					l = append(l, stringedCell, stringedCell)
				}
			}
			field = append(field, l)
		}
	}
	return field, instructions

}

func renderField(field [][]string) {
	output := "\n\n"
	for _, line := range field {
		output += strings.Join(line, "") + "\n"
	}
	log.Println(output)
}

func calculateGPS(field [][]string) int {
	total := 0
	for y, line := range field {
		for x, cell := range line {
			if cell == "O" {
				total += 100*y + x
			}
		}
	}
	return total
}

func canMove(x, y int, instruction string, field [][]string) bool {
	dx, dy := 0, 0
	switch instruction {
	case "^":
		dy = -1
	case "v":
		dy = 1
	case "<":
		dx = -1
	case ">":
		dx = 1
	}
	if field[y+dy][x+dx] == "." {
		return true
	}
	if field[y+dy][x+dx] == "#" {
		return false
	}
	if field[y+dy][x+dx] == "O" {
		canMove := canMove(x+dx, y+dy, instruction, field)
		return canMove
	}
	if dy == 0 && (field[y+dy][x+dx] == "[" || field[y+dy][x+dx] == "]") {
		canMove := canMove(x+dx, y+dy, instruction, field)
		return canMove
	}
	return false
}

func move(x, y int, instruction string, field [][]string) {
	dx, dy := 0, 0
	switch instruction {
	case "^":
		dy = -1
	case "v":
		dy = 1
	case "<":
		dx = -1
	case ">":
		dx = 1
	}
	log.Println("move", x, y, field[y][x], instruction)
	for true {
		newX := x + dx
		newY := y + dy
		if field[newY][newX] == "." {
		}
	}
	previous := field[y][x]
	if field[y+dy][x+dx] == "." {
		field[y+dy][x+dx] = previous
		field[y][x] = "."
		if previous == "@" {
			robot.x += dx
			robot.y += dy
		}
	} else {
		move(x+dx, y+dy, instruction, field)
	}
}

func task1(field [][]string, instructions []string) int {
	for _, instruction := range instructions {
		if canMove(robot.x, robot.y, instruction, field) {
			move(robot.x, robot.y, instruction, field)
		}
	}
	renderField(field)
	return calculateGPS(field)
}

func task2(field [][]string, instructions []string) int {
	for _, instruction := range instructions {
		if canMove(robot.x, robot.y, instruction, field) {
			renderField(field)
			move(robot.x, robot.y, instruction, field)
		}
	}
	return 0
}
func main() {
	file, err := os.ReadFile("2024/15/test2.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	input := string(file)
	field, instructions := parseInputWide(input)
	// log.Println("task1: ", task1(field, instructions))
	log.Println("task2: ", task2(field, instructions))
}
