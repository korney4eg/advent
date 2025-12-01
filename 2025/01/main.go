package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile(os.Args[1])
	input := string(f)

	dial := make([]int, 100)

	// fmt.Println(input)
	dials := parseInput(input)
	_ = process(dial, 50, dials)
	_ = process2(dial, 50, dials)
}

func parseInput(input string) (output []int) {
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		direction := string(line[0])
		stepsS := line[1:len(line)]
		steps, _ := strconv.Atoi(stepsS)
		if direction == "L" {
			output = append(output, -1*steps)
		} else {
			output = append(output, steps)
		}

	}
	// fmt.Println(output)
	return output
}

func process(dial []int, start int, turns []int) (positions []int) {
	current := start
	prev := start
	password := 0
	for _, turn := range turns {
		// fmt.Println(current, turn)
		prev = current
		if math.Abs(float64(turn)) > float64(len(dial)) {
			turn = turn % (len(dial))
		}
		if current+turn < 0 {
			current = len(dial) + (turn + current)
		} else if current+turn >= len(dial) {
			current = current + turn - len(dial)
		} else {
			current = current + turn
		}
		if current > len(dial) || current < 0 {
			fmt.Println("current", current, "prev", prev, "turn", turn)
		}
		positions = append(positions, current)
		if current == 0 {
			password++
		}

	}
	fmt.Println("password", password)
	return positions
}

func process2(dial []int, start int, turns []int) (positions []int) {
	current := start
	password := 0
	for _, turn := range turns {
		miniturn := turn % len(dial)
		steps := []int{miniturn}
		move := 0
		if turn > 0 {
			move = 1
		} else if turn < 0 {
			move = -1
		}
		for j := 0; math.Abs(float64(j)) < math.Abs(float64(turn/len(dial))); j += move {

			steps = append(steps, len(dial))
		}
		fmt.Println(turn, steps, turn/len(dial))
		for _, step := range steps {
			changes := 0
			next := 0
			// fmt.Println(current, turn)
			fmt.Println("current", current, "+", step)
			// if turn > 0 && turn >= len(dial) {
			// 	fmt.Println("got more")
			// 	changes += turn / len(dial)
			// } else if turn < 0 && -1*turn >= len(dial) {
			// 	fmt.Println("got less")
			// }
			if current+step < 0 {
				if current != 0 {
					fmt.Println("click -")
					changes++
				}
				next = len(dial) + (step + current)
			} else if current+step >= len(dial) {
				next = current + step - len(dial)
				if next != 0 {
					fmt.Println("click +")
					changes++
				}
			} else {
				next = current + step
			}
			// if current > len(dial) || current < 0 {
			// 	fmt.Println("current", current, "prev", prev, "step", step)
			// }
			positions = append(positions, next)
			if next == 0 {
				password++
				fmt.Println("click 0")
			}
			current = next
			password += changes

		}
	}
	fmt.Println("password", password)
	return positions
}
