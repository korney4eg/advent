package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// var maxX, maxY = 11, 7

var maxX, maxY = 101, 103

type robot struct {
	x, y   int
	vX, vY int
}

func readInput(input string) (robots []*robot) {
	for _, line := range strings.Split(input, "\n") {
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		r := &robot{}
		position := strings.TrimPrefix(strings.Split(line, " ")[0], "p=")
		xs := strings.Split(position, ",")[0]
		ys := strings.Split(position, ",")[1]
		r.x, _ = strconv.Atoi(xs)
		r.y, _ = strconv.Atoi(ys)

		velocity := strings.TrimPrefix(strings.Split(line, " ")[1], "v=")
		vxs := strings.Split(velocity, ",")[0]
		vys := strings.Split(velocity, ",")[1]
		r.vX, _ = strconv.Atoi(vxs)
		r.vY, _ = strconv.Atoi(vys)
		robots = append(robots, r)
	}
	return robots
}

func (r *robot) move(seconds int) {
	newX := r.x + r.vX*seconds
	newX = newX % maxX
	if newX < 0 {
		newX += maxX
	}
	r.x = newX
	newY := r.y + r.vY*seconds
	newY = newY % maxY
	if newY < 0 {
		newY += maxY
	}
	r.y = newY
}

var field [][]int = make([][]int, maxY)

func render(robots []*robot, showDelim bool) {
	for y := 0; y < maxY; y++ {
		line := make([]int, maxX)
		for x := 0; x < maxX; x++ {
			line[x] = 0

		}
		field[y] = line
	}
	for _, r := range robots {
		field[r.y][r.x]++
	}
	output := ""
	for y := 0; y < maxY; y++ {
		output += "\n"
		for x := 0; x < maxX; x++ {
			if showDelim && (x == maxX/2 || y == maxY/2) {
				output += " "
			} else if field[y][x] == 0 {
				output += " "

			} else {
				output += "#"
				// output += strconv.Itoa(field[y][x])
			}
		}
	}

	if strings.Contains(output, "######") {
		fmt.Println(output)
	}
}

func calculate(robots []*robot) int {
	result := 1
	counter := make(map[int]int)
	for _, r := range robots {
		if r.x < maxX/2 && r.y < maxY/2 {
			counter[0]++
		} else if r.x < maxX/2 && r.y > maxY/2 {
			counter[2]++

		}
		if r.x > maxX/2 && r.y < maxY/2 {
			counter[1]++
		} else if r.x > maxX/2 && r.y > maxY/2 {
			counter[3]++

		}

	}
	log.Println(counter)
	for _, v := range counter {
		result *= v
	}

	return result
}

func task1(input string) int {
	robots := readInput(input)

	render(robots, true)
	for _, r := range robots {
		r.move(100)
	}
	render(robots, true)

	return calculate(robots)
}
func task2(input string) int {
	robots := readInput(input)

	render(robots, false)
	for i := 1; i < 100000; i++ {
		for _, r := range robots {
			r.move(1)
		}
		fmt.Println(i)
		render(robots, false)
	}

	return 0
}

func main() {
	file, err := os.ReadFile("2024/14/input.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	input := string(file)

	// log.Printf("Task 1: %d", task1(input))
	log.Printf("Task 2: %d", task2(input))
}
