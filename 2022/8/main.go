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

	visibleTreesCount := 0

	for rowIndex := 0; rowIndex < len(forest); rowIndex++ {
		row := forest[rowIndex]

		if rowIndex == 0 || (rowIndex == len(forest)-1) {
			visibleTreesCount += len(row)
			continue
		}

		for columnIndex := 0; columnIndex < len(row); columnIndex++ {
			if columnIndex == 0 || (columnIndex == len(row)-1) {
				visibleTreesCount++
				continue
			}

			tree := Tree{
				rowIndex:    rowIndex,
				columnIndex: columnIndex,
				height:      row[columnIndex],
			}

			if !getIsHiddenTree(tree, forest) {
				visibleTreesCount++
			}
		}

	}

	fmt.Printf("%v", forest)
	fmt.Printf("%v", visibleTreesCount)
}

func getIsHiddenTree(tree Tree, forest [][]int) bool {
	treeHeight := tree.height
	treeRowIndex := tree.rowIndex
	treeColumnIndex := tree.columnIndex

	isTopCover := false
	isRightCover := false
	isBottomCover := false
	isLeftCover := false

	for rowIndex := 0; rowIndex < len(forest); rowIndex++ {
		row := forest[rowIndex]

		for columnIndex := 0; columnIndex < len(row); columnIndex++ {
			isCurrentTreeRow := rowIndex == treeRowIndex
			isCurrentTreeColumn := columnIndex == treeColumnIndex
			isCurrentTree := isCurrentTreeRow && isCurrentTreeColumn

			// skip current tree
			if isCurrentTree {
				continue
			}

			// skip non tree row AND column
			if !isCurrentTreeRow && !isCurrentTreeColumn {
				continue
			}

			candidateTreeHeight := forest[rowIndex][columnIndex]

			if !isTopCover {
				isTopCover = rowIndex < treeRowIndex && candidateTreeHeight >= treeHeight
			}
			if !isBottomCover {
				isBottomCover = rowIndex > treeRowIndex && candidateTreeHeight >= treeHeight
			}
			if !isLeftCover {
				isLeftCover = columnIndex < treeColumnIndex && candidateTreeHeight >= treeHeight
			}
			if !isRightCover {
				isRightCover = columnIndex > treeColumnIndex && candidateTreeHeight >= treeHeight
			}
		}
	}

	return isTopCover && isRightCover && isBottomCover && isLeftCover
}
