package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type button struct {
	x, y int
}

type point struct {
	x, y int
}

type claw struct {
	a, b  button
	prize point
}

var aCost = 3
var bCost = 1

func (c claw) calculate(initCoord int) (a, b int, found bool) {
	prizeX := c.prize.x + initCoord
	prizeY := c.prize.y + initCoord
	minButton := button{}
	maxButton := button{}
	if (c.a.x/c.b.x)*(c.a.y/c.b.y) > 1 {
		maxButton = c.a
		minButton = c.b
	} else {
		maxButton = c.b
		minButton = c.a
	}
	log.Printf("commonX for %d and %d = %d:", c.b.x, c.a.x, commonX)
	log.Printf("commonY for %d and %d = %d:", c.b.y, c.a.y, commonY)

	maxB := int(math.Max(float64(prizeX/commonX), float64(prizeY/commonY)))
	// if maxB > 180 {
	// 	maxB = 180
	// }
	for i := maxB - 1; i > 0; i-- {
		totalBx := c.b.x * i
		totalBy := c.b.y * i
		j := (prizeX - totalBx) / c.a.x
		totalAx := c.a.x * j
		totalAy := c.a.y * j

		if prizeX-(totalBx+totalAx) == 0 && prizeY-(totalBy+totalAy) == 0 {
			return j, i, true
		}
	}

	return 0, 0, false

}

func commonMinimum(a, b int) int {
	min := 99999999
	max := 0
	if a < b {
		max = b
		min = a
	} else {
		max = a
		min = b
	}
	// if min%2 == 0 && max%2 == 0 {
	// 	min = min / 2
	// 	max = max / 2
	// }
	log.Println("min:", min, "max:", max)
	for i := min; i <= min*max; i += min {
		if i%a == 0 && i%b == 0 {
			log.Println("Common minimum:", i)
			return i
		}
	}
	return 0
}

func task1(input string) int {
	all := 0
	claws := read(input)
	for _, claw := range claws {
		a, b, found := claw.calculate(0)
		if found {
			log.Printf("%+v: %d * %d + %d + %d", claw, a, aCost, b, bCost)
			all += a*aCost + b*bCost
		} else {
			log.Printf("%+v skipped", claw)
		}
	}
	return all
}
func task2(input string) int {
	all := 0
	claws := read(input)
	for _, claw := range claws {
		a, b, found := claw.calculate(10000000000000)
		if found {
			log.Printf("%+v: %d * %d + %d + %d", claw, a, aCost, b, bCost)
			all += a*aCost + b*bCost
		} else {
			log.Printf("%+v skipped", claw)
		}
	}
	return all
}
func read(input string) []claw {
	var claws []claw
	c := claw{}
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "Button A") {
			Axs := strings.Split(strings.Split(line, ":")[1], ",")[0][3:]
			Ays := strings.Split(strings.Split(line, ":")[1], ",")[1][3:]
			Abutton := button{}
			Abutton.x, _ = strconv.Atoi(Axs)
			Abutton.y, _ = strconv.Atoi(Ays)
			c.a = Abutton
		} else if strings.HasPrefix(line, "Button B") {
			Bxs := strings.Split(strings.Split(line, ":")[1], ",")[0][3:]
			Bys := strings.Split(strings.Split(line, ":")[1], ",")[1][3:]
			Bbutton := button{}
			Bbutton.x, _ = strconv.Atoi(Bxs)
			Bbutton.y, _ = strconv.Atoi(Bys)
			c.b = Bbutton
		} else if strings.HasPrefix(line, "Prize:") {
			prize := point{}
			Pxs := strings.Split(strings.Split(line, ":")[1], ",")[0][3:]
			Pys := strings.Split(strings.Split(line, ":")[1], ",")[1][3:]
			prize.x, _ = strconv.Atoi(Pxs)
			prize.y, _ = strconv.Atoi(Pys)
			c.prize = prize
		} else {
			claws = append(claws, c)
			c = claw{}
		}
	}
	log.Printf("Claws: %+v", claws)
	return claws
}

func main() {
	file, err := os.ReadFile("2024/13/input.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	input := string(file)
	input = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

	log.Printf("Task 1: %d", task1(input))
	// log.Printf("Task 2: %d", task2(input))

}
