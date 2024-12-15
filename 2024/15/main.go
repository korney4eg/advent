package main

import (
	"log"
	"os"
	"strings"
)

type position struct {
	x, y int
}

var field = [][]string{}
var instructions = []string{}
var robot = position{}

func parseInput(input string) {
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

}

func renderField() {
	output := "\n\n"
	for _, line := range field {
		output += strings.Join(line, "") + "\n"
	}
	log.Println(output)
}

func calculateGPS() int {
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

func move(x, y int, instruction string) bool {
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
	log.Println("move", x, y, instruction)
	previous := field[y][x]
	if field[y+dy][x+dx] == "." {
		field[y][x] = "."
		field[y+dy][x+dx] = previous
		if previous == "@" {
			robot.x += dx
			robot.y += dy
		}
		return true
	}
	if field[y+dy][x+dx] == "#" {
		return false
	}
	if field[y+dy][x+dx] == "O" {
		canMove := move(x+dx, y+dy, instruction)
		if canMove {
			field[y][x] = "."
			field[y+dy][x+dx] = previous
			if previous == "@" {
				robot.x += dx
				robot.y += dy
			}
		}
		return canMove
	}

	return false
}

func task1() int {
	for _, instruction := range instructions {
		_ = move(robot.x, robot.y, instruction)
	}
	renderField()
	return calculateGPS()
}

func main() {
	file, err := os.ReadFile("2024/15/input.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	input := string(file)
	_ = `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`
	parseInput(input)
	log.Println("task1: ", task1())
}
