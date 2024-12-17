package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

type field struct {
	rendered [][]string
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
	for _, line := range strings.Split(f.field, "\n") {
		if len(line) == 0 {
			continue
		}
		l := []string{}
		f.board = append(f.board, strings.Split(line, ""))
		for _ = range line {
			l = append(l, ".")
		}
		f.rendered = append(f.rendered, l)
	}
}

func (f *field) render() {
	for _, line := range f.rendered {
		fmt.Println(strings.Join(line, ""))
	}
	fmt.Println()
	fmt.Println()
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

func (f *field) hasXmax(x, y, dx, dy int, regex string) bool {

	line := ""
	if x+(len(regex)-1)*dx < 0 || x+(len(regex)-1)*dx >= len(f.board[0]) || y+(len(regex)-1)*dy < 0 || y+(len(regex)-1)*dy >= len(f.board) {
		return false
	}
	founded := [][2]int{}
	for i := 0; i < len(regex); i++ {
		line += f.board[y+i*dy][x+i*dx]
		founded = append(founded, [2]int{y + i*dy, x + i*dx})
	}
	if line == regex {
		for i := 0; i < len(regex); i++ {
			f.rendered[y+i*dy][x+i*dx] = f.board[y+i*dy][x+i*dx]
		}
	}
	return line == regex
}

func (f *field) findXMAS(lines []string, Xs [][2]int) int {
	count := 0
	for _, X := range Xs {
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				if dx == 0 && dy == 0 {
					continue
				}
				// log.Printf("X: %d %d", X[0], X[1])
				if f.hasXmax(X[1], X[0], dx, dy, `XMAS`) {
					count++
				}
			}
		}
	}
	return count
}

func (f *field) findMAS(lines []string, As [][2]int) int {
	count := 0
	for _, A := range As {
		// log.Printf("X: %d %d", X[0], X[1])
		if A[0] == 0 || A[0] == len(lines)-1 || A[1] == 0 || A[1] == len(lines[0])-1 {
			continue
		}
		line1 := []string{string(lines[A[0]-1][A[1]-1]), string(lines[A[0]][A[1]]), string(lines[A[0]+1][A[1]+1])}
		line2 := []string{string(lines[A[0]+1][A[1]-1]), string(lines[A[0]][A[1]]), string(lines[A[0]-1][A[1]+1])}
		sort.Strings(line1)
		sort.Strings(line2)
		if strings.Join(line1, "") == "AMS" && strings.Join(line2, "") == "AMS" {
			f.rendered[A[0]][A[1]] = f.board[A[0]][A[1]]
			f.rendered[A[0]-1][A[1]-1] = f.board[A[0]-1][A[1]-1]
			f.rendered[A[0]+1][A[1]+1] = f.board[A[0]+1][A[1]+1]
			f.rendered[A[0]+1][A[1]-1] = f.board[A[0]+1][A[1]-1]
			f.rendered[A[0]-1][A[1]+1] = f.board[A[0]-1][A[1]+1]
			// f.render()

			count++
		}
	}
	return count
}

func (f *field) task1() int {
	count := 0
	log.Printf("Field:\n%s", f.field)
	lines := strings.Split(f.field, "\n")
	// find Xs
	Xs := [][2]int{}
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] == 'X' {
				Xs = append(Xs, [2]int{y, x})
			}
		}
	}
	count = f.findXMAS(lines, Xs)
	return count
}

func (f *field) task2() int {
	count := 0
	log.Printf("Field:\n%s", f.field)
	lines := strings.Split(f.field, "\n")
	// find Xs
	As := [][2]int{}
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] == 'A' {
				As = append(As, [2]int{y, x})
			}
		}
	}
	count = f.findMAS(lines, As)
	return count
}

func main() {
	file, err := os.ReadFile("2024/04/input.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	inputs := string(file)
	f := field{}
	f.read(inputs)
	count := f.task2()
	log.Println(count)
}
