package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tree struct {
	rowIndex    int
	columnIndex int
	height      int
}

/*
row ->
30373
25512
65332
33549
35390
*/

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var forest [][]int
	var linesList []string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		linesList = append(linesList, line)
	}

	for _, line := range linesList {
		var newLine []int
		forestLine := strings.Split(line, "")
		for _, char := range forestLine {
			tree, _ := strconv.Atoi(char)
			newLine = append(newLine, tree)
		}
		forest = append(forest, newLine)
	}

	threeScore := 0

	for rowIndex := 0; rowIndex < len(forest); rowIndex++ {
		row := forest[rowIndex]

		for columnIndex := 0; columnIndex < len(row); columnIndex++ {
			tree := Tree{
				rowIndex:    rowIndex,
				columnIndex: columnIndex,
				height:      row[columnIndex],
			}

			newTreeScore := getTreeScore(tree, forest)

			if newTreeScore > threeScore {
				threeScore = newTreeScore
			}

		}

	}

	fmt.Println(threeScore)
}

func getTreeScore(tree Tree, forest [][]int) int {
	treeHeight := tree.height
	treeRowIndex := tree.rowIndex
	treeColumnIndex := tree.columnIndex

	isTopFinished := false
	topScore := 0

	isRightFinished := false
	rightScore := 0

	isBottomFinished := false
	bottomScore := 0

	isLeftFinished := false
	leftScore := 0

	// move down
	for rowIndex := treeRowIndex + 1; rowIndex < len(forest); rowIndex++ {
		if isBottomFinished {
			break
		}
		bottomScore++

		isBottomFinished = forest[rowIndex][treeColumnIndex] >= treeHeight
	}

	// move up
	for rowIndex := treeRowIndex - 1; rowIndex >= 0; rowIndex-- {
		if isTopFinished {
			break
		}
		topScore++

		isTopFinished = forest[rowIndex][treeColumnIndex] >= treeHeight
	}

	row := forest[treeRowIndex]

	// move right
	for columnIndex := treeColumnIndex + 1; columnIndex < len(row); columnIndex++ {
		if isRightFinished {
			break
		}
		rightScore++

		isRightFinished = row[columnIndex] >= treeHeight
	}

	// move left
	for columnIndex := treeColumnIndex - 1; columnIndex >= 0; columnIndex-- {
		if isLeftFinished {
			break
		}
		leftScore++

		isLeftFinished = row[columnIndex] >= treeHeight
	}

	return topScore * rightScore * bottomScore * leftScore
}
