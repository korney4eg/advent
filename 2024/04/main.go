package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type field struct {
	rendered string
	field    string
	board    [][]string
}

func rotate(input string) string {
	lines := strings.Split(input, "\n")
	rows := len(lines)
	cols := len(lines[0])
	// 90 degrees clockwise
	rotated := make([]string, cols)
	for i := range rotated {
		rotated[i] = ""
	}
	for col := 0; col < cols; col++ {
		for row := rows - 1; row >= 0; row-- {
			rotated[col] += string(lines[row][col])
		}
	}
	return strings.Join(rotated, "\n")
}

func countInDiagonals(field string) int {
	count := 0
	urRe := regexp.MustCompile(`X...\n.M..\n..A.\n...S`)
	if urRe.Match([]byte(field)) {
		count++
	}
	dlRe := regexp.MustCompile(`...S\n..A.\n.M..\nX...`)
	if dlRe.Match([]byte(field)) {
		count++
	}
	ulRe := regexp.MustCompile(`...X\n..M.\n.A..\nS...`)
	if ulRe.Match([]byte(field)) {
		count++
	}
	drRe := regexp.MustCompile(`S...\n.A..\n..M.\n...X`)
	if drRe.Match([]byte(field)) {
		count++
	}
	return count
}

func findInField(field string) int {
	count := 0
	count += strings.Count(field, "XMAS")
	count += strings.Count(field, "SAMX")
	return count
}

func (f *field) read(input string) {
	f.field = strings.TrimRight(input, "\n")
}

func (f *field) render() {
	for _, line := range f.board {
		fmt.Println(strings.Join(line, ""))
	}
}

func (f *field) find() int {
	count := 0
	log.Printf("Field:\n%s", f.field)
	lines := strings.Split(f.field, "\n")
	for y := 0; y < len(lines)-3; y++ {
		for x := 0; x < len(lines[y])-3; x++ {
			newField := strings.Join([]string{lines[y][x : x+4], lines[y+1][x : x+4], lines[y+2][x : x+4], lines[y+3][x : x+4]}, "\n")
			// boardedField := strings.Split(newField, "\n")
			// log.Println(boardedField)
			log.Printf("small:\n%s", newField)
			count += findInField(newField)
			log.Printf("normal: %d", findInField(newField))
			count += countInDiagonals(newField)
			log.Printf("diags: %d", countInDiagonals(newField))
			count += countInDiagonals(rotate(newField))
			log.Printf("reverted: %d", countInDiagonals(rotate(newField)))
		}
	}
	return count
}

func findXMAS(lines []string, Xs [][]int) int {
	count := 0
	newLines := []string{}
	re := regexp.MustCompile(`.`)
	for _, line := range lines {
		newLines = append(newLines, re.ReplaceAll(`.`, `.`))
	}
	for _, X := range Xs {
		y, x := X[0], X[1]
		canLeft := x-4 >= 0
		canRight := x+4 < len(lines[y])
		canUp := y-4 >= 0
		canDown := y+4 < len(lines)
		if canRight && lines[y][x:x+4] == "XMAS" {
			count++
		}
		if canLeft && lines[y][x-4:x] == "SAMX" {
			count++
		}
		if canDown && (string(lines[y][x]) == "X" && string(lines[y+1][x]) == "M" && string(lines[y+2][x]) == "A" && string(lines[y+3][x]) == "S") {
			count++
		}
		if canUp && (string(lines[y][x]) == "X" && string(lines[y-1][x]) == "M" && string(lines[y-2][x]) == "A" && string(lines[y-3][x]) == "S") {
			count++
		}
		if canUp && canRight && (string(lines[y][x]) == "X" && string(lines[y-1][x+1]) == "M" && string(lines[y-2][x+2]) == "A" && string(lines[y-3][x+3]) == "S") {
			count++
		}
		if canDown && canRight && (string(lines[y][x]) == "X" && string(lines[y+1][x+1]) == "M" && string(lines[y+2][x+2]) == "A" && string(lines[y+3][x+3]) == "S") {
			count++
		}
		if canUp && canLeft && (string(lines[y][x]) == "X" && string(lines[y-1][x-1]) == "M" && string(lines[y-2][x-2]) == "A" && string(lines[y-3][x-3]) == "S") {
			count++
		}
		if canDown && canLeft && (string(lines[y][x]) == "X" && string(lines[y+1][x-1]) == "M" && string(lines[y+2][x-2]) == "A" && string(lines[y+3][x-3]) == "S") {
			count++
		}
	}
	return count
}

func (f *field) find2() int {
	count := 0
	log.Printf("Field:\n%s", f.field)
	lines := strings.Split(f.field, "\n")
	// find Xs
	Xs := [][]int{}
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] == 'X' {
				Xs = append(Xs, []int{y, x})
			}
		}
	}
	count = findXMAS(lines, Xs)
	return count
}

func main() {
	file, err := os.ReadFile("test.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	inputs := string(file)
	f := field{}
	f.read(inputs)
	count := f.find2()
	log.Println(count)
}
