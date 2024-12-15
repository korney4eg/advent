package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/fatih/color"
)

type point struct {
	x, y int
}

var field [][]string
var visited [][]bool

func render(pieces []point, gotX, gotY int) {
	output := "\n"
	for y := 0; y < len(field); y++ {
		line := ""
		for x := 0; x < len(field[y]); x++ {
			if x == gotX && y == gotY {
				line += color.RedString(fmt.Sprintf("%v", field[y][x]))
			} else if slices.Contains(pieces, point{x, y}) {
				line += color.GreenString(fmt.Sprintf("%v", field[y][x]))
			} else {
				line += fmt.Sprintf("%v", field[y][x])
			}
		}
		output += line + "\n"
	}
	log.Println(output)
}

func findAllNeighbours(x, y int) (result []point, fences int) {
	allFeneces := 0
	toVisit := []point{{x, y}}
	for len(toVisit) > 0 {
		symbol := field[toVisit[0].y][toVisit[0].x]
		fences = 0
		toVisitX := toVisit[0].x
		toVisitY := toVisit[0].y
		if visited[toVisitY][toVisitX] {
			toVisit = slices.Delete(toVisit, 0, 1)
			continue
		}
		visited[toVisitY][toVisitX] = true
		if !slices.Contains(result, point{toVisitX, toVisitY}) {
			result = append(result, point{toVisitX, toVisitY})
		}
		for _, nextStep := range []point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			newX := toVisitX + nextStep.x
			newY := toVisitY + nextStep.y
			if slices.Contains(result, point{newX, newY}) {
				continue
			}
			log.Println("newX:", newX, "newY:", newY)
			if (newX >= 0 && newX < len(field[0]) && newY >= 0 && newY < len(field)) &&
				field[newY][newX] == symbol &&
				!slices.Contains(result, point{newX, newY}) {
				toVisit = append(toVisit, point{newX, newY})
				continue
			}
			fences++

		}
		toVisit = slices.Delete(toVisit, 0, 1)
		// 	if field[toVisitY][toVisitX] == "T" {
		// 		render(result, toVisitX, toVisitY)
		// 		log.Println("X:", toVisitX, "Y:", toVisitY)
		// 		log.Println("fences:", fences, "area:", len(result))
		// 	}
		allFeneces += fences
	}
	return result, allFeneces

}

func task1(input string) int {
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		visitedLine := []bool{}
		for _, _ = range line {
			visitedLine = append(visitedLine, false)

		}
		field = append(field, strings.Split(line, ""))
		visited = append(visited, visitedLine)
	}
	log.Println(field)
	// var pieces [][]point
	all := 0
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			if visited[y][x] {
				continue
			}
			area, fences := findAllNeighbours(x, y)
			log.Println("A region of ", field[y][x], "plants with price ", fences, "*", len(area), "=", fences*len(area))
			all += fences * len(area)

		}
	}
	return all
}

func main() {
	file, err := os.ReadFile("2024/12/input.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	input := string(file)
	// 	input = `RRRRIICCFF
	// RRRRIICCCF
	// VVRRRCCFFF
	// VVRCCCJFFF
	// VVVVCJJCFE
	// VVIVCCJJEE
	// VVIIICJJEE
	// MIIIIIJJEE
	// MIIISIJEEE
	// MMMISSJEEE`
	log.Printf("Task 1: %d", task1(input))
}
