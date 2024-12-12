package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var board []int
var startNum = 0
var finishNum = 9
var startPositions []int
var cols, rows int
var visitedPath [][]int

func task1(input string) int {
	readBoard(input)
	for _, pos := range startPositions {
		path, success := bfs(pos)
		if success {
			return len(path)
		}
	}
	return 0
}

func positionToXY(position int) (x, y int) {
	x = position % cols
	y = position / cols
	return x, y

}

func getNextPositions(position int) (nextPositions []int) {
	x, y := positionToXY(position)

	if x-1 >= 0 && (board[position-1] == board[position]+1 || board[position-1] == board[position]-1) {
		nextPositions = append(nextPositions, position-1)
	}
	if x+1 < cols && (board[position+1] == board[position]+1 || board[position+1] == board[position]-1) {
		nextPositions = append(nextPositions, position+1)
	}
	if y-1 >= 0 && (board[position-rows] == board[position]+1 || board[position-rows] == board[position]-1) {
		nextPositions = append(nextPositions, position-rows)
	}
	if y+1 < rows && (board[position+rows] == board[position]+1 || board[position+rows] == board[position]-1) {
		nextPositions = append(nextPositions, position+rows)
	}
	return nextPositions

}

var iterations = 0

func bfs(position int) (path []int, success bool) {
	x, y := positionToXY(position)
	log.Println("position", position, x, y)
	if board[position] == finishNum {
		return []int{position}, true
	}
	nextSteps := getNextPositions(position)
	if len(nextSteps) == 0 {
		return []int{}, false
	}
	for _, nextStep := range nextSteps {
		newPath := []int{position, nextStep}
		if slices.Contains(visitedPath, newPath) {
			continue
		}
	}
	log.Println("nextSteps", nextSteps)
	return path, success
}

func readBoard(input string) {
	visitedPath = make([][]int, 0)
	cols = strings.Index(input, "\n")
	rows = len(strings.Split(input, "\n")) - 1
	for i, c := range strings.Split(strings.ReplaceAll(input, "\n", ""), "") {
		if c == "" {
			continue
		}
		num, _ := strconv.Atoi(c)
		if num == 0 {
			startPositions = append(startPositions, i)
		}
		board = append(board, num)
	}
	log.Println("board", board)
	log.Println("maxX", cols, "maxY", rows)
}

func main() {
	file, err := os.ReadFile("2024/10/test1.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	input := string(file)
	// input = `2333133121414131402`
	log.Println("task2: ", task1(input))
}
