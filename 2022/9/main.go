package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const ropeLen = 10

type knot struct {
	previousKnot *knot
	x            int
	y            int
	oldX         int
	oldY         int
}

func (b *knot) makeStep(direction string) {
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

func (b *knot) makeMoveToPreviousPart() {
	if b.previousKnot == nil {
		return
	}
	dx := b.previousKnot.x - b.x
	dy := b.previousKnot.y - b.y
	if math.Abs(float64(dx))+math.Abs(float64(dy)) > 2 {
		b.x += int(dx / int(math.Abs(float64(dx))))
		b.y += int(dy / int(math.Abs(float64(dy))))

	} else if b.x < b.previousKnot.x-1 {
		b.x++
	} else if b.x > b.previousKnot.x+1 {
		b.x--
	} else if b.y < b.previousKnot.y-1 {
		b.y++
	} else if b.y > b.previousKnot.y+1 {
		b.y--
	}
}

type field struct {
	maxX    int
	maxY    int
	minX    int
	minY    int
	objects []*knot
}

func (f *field) draw() {

	for i := f.maxY; i >= f.minY; i-- {
		newLine := ""
		for j := f.minX; j <= f.maxX; j++ {
			newSymbol := "."
			if i == j && i == 0 {
				newSymbol = "s"
			}
			for objIndex, obj := range f.objects {
				// fmt.Printf("newline:'%s',j=%d,i=%d,getting object %d - %d:%d\n", newLine, j, i, objIndex, obj.x, obj.y)
				if obj.x == j && obj.y == i && (newSymbol == "." || newSymbol == "s") {
					newSymbol = fmt.Sprintf("%d", objIndex)
				}

			}
			newLine += newSymbol
		}
		fmt.Println(newLine)
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

	newField := field{
		maxX: 13,
		maxY: 20,
		minX: -10,
		minY: -10,
	}
	input, _ := os.ReadFile("input.txt")
	head := &knot{x: 0, y: 0}
	rope := []*knot{head}
	for i := 1; i < ropeLen; i++ {
		newKnot := &knot{x: 0, y: 0, previousKnot: rope[i-1]}
		rope = append(rope, newKnot)
	}
	newField.objects = rope
	tailWasHere := []*coord{}
	newField.draw()

	for _, mv := range parseMoves(string(input)) {
		fmt.Printf("Move: %s %d\n", mv.direction, mv.steps)
		for step := 0; step < mv.steps; step++ {
			head.makeStep(mv.direction)
			fmt.Printf("head %d:%d\n", head.x, head.y)
			for i := 1; i < ropeLen; i++ {
				rope[i].makeMoveToPreviousPart()
			}
			tail := rope[len(rope)-1]
			coordinate := &coord{x: tail.x, y: tail.y}
			tailWasHere = addIfNot(tailWasHere, coordinate)

			fmt.Println()
		}

	}
	fmt.Println(len(tailWasHere))

}
