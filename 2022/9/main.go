package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bodyPart struct {
	previousPart *bodyPart
	nextPart     *bodyPart
	x            int
	y            int
	oldX         int
	oldY         int
}

func (b *bodyPart) makeStep(direction string) {
	var dx, dy int
	switch direction {
	case "R":
		dx++
	case "L":
		dx--
	case "U":
		dy++
	case "D":
		dy--
	}
	b.oldX = b.x
	b.oldY = b.y
	b.x += dx
	b.y += dy
}

func (b *bodyPart) makeMoveToPreviousPart() {
	if b.previousPart == nil {
		return
	}
	if b.x < b.previousPart.x-1 || b.x > b.previousPart.x+1 || b.y < b.previousPart.y-1 || b.y > b.previousPart.y+1 {
		b.x = b.previousPart.oldX
		b.y = b.previousPart.oldY
	}
}

func parseMoves(moveSet string) (moves []move) {
	movesStringList := strings.Split(moveSet, "\n")
	for _, moveString := range movesStringList {
		if moveString == "" {
			continue
		}
		direction := strings.Split(moveString, " ")[0]
		steps, _ := strconv.Atoi(strings.Split(moveString, " ")[1])
		mv := move{direction: direction, steps: steps}
		moves = append(moves, mv)
	}
	return moves
}

type move struct {
	direction string
	steps     int
}

type coord struct {
	x int
	y int
}

func addIfNot(coords []*coord, coordinate *coord) (updatedCoords []*coord) {
	for _, listedCoord := range coords {
		if listedCoord.x == coordinate.x && listedCoord.y == coordinate.y {
			return coords
		}
	}
	return append(coords, coordinate)
}

func main() {

	input, _ := os.ReadFile("input.txt")
	head := &bodyPart{x: 0, y: 0}
	tail := &bodyPart{x: 0, y: 0, previousPart: head}
	head.nextPart = tail
	tailWasHere := []*coord{}

	for _, mv := range parseMoves(string(input)) {
		fmt.Printf("Move: %s %d\n", mv.direction, mv.steps)
		for step := 0; step < mv.steps; step++ {
			head.makeStep(mv.direction)
			tail.makeMoveToPreviousPart()
			fmt.Printf("head:%d,%d; tail: %d:%d\n", head.x, head.y, tail.x, tail.y)
			coordinate := &coord{x: tail.x, y: tail.y}
			tailWasHere = addIfNot(tailWasHere, coordinate)
		}

	}
	fmt.Println(len(tailWasHere))

}
