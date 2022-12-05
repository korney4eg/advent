package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// input := `    [D]
	// [N] [C]
	// [Z] [M] [P]`
	// commandsText := `move 1 from 2 to 1
	// move 3 from 1 to 3
	// move 2 from 2 to 1
	// move 1 from 1 to 2`

	input := `    [V] [G]             [H]        
[Z] [H] [Z]         [T] [S]        
[P] [D] [F]         [B] [V] [Q]    
[B] [M] [V] [N]     [F] [D] [N]    
[Q] [Q] [D] [F]     [Z] [Z] [P] [M]
[M] [Z] [R] [D] [Q] [V] [T] [F] [R]
[D] [L] [H] [G] [F] [Q] [M] [G] [W]
[N] [C] [Q] [H] [N] [D] [Q] [M] [B]`

	commandsText := `move 3 from 2 to 5
	move 2 from 9 to 6
	move 4 from 7 to 1
	move 7 from 3 to 4
	move 2 from 9 to 8
	move 8 from 8 to 6
	move 1 from 7 to 4
	move 8 from 6 to 4
	move 4 from 5 to 7
	move 3 from 4 to 9
	move 2 from 6 to 3
	move 11 from 4 to 1
	move 1 from 3 to 4
	move 2 from 3 to 1
	move 1 from 7 to 6
	move 14 from 1 to 6
	move 7 from 4 to 3
	move 2 from 5 to 9
	move 5 from 6 to 4
	move 9 from 6 to 1
	move 3 from 4 to 8
	move 1 from 7 to 6
	move 3 from 4 to 1
	move 7 from 3 to 8
	move 5 from 9 to 5
	move 4 from 1 to 4
	move 3 from 7 to 2
	move 5 from 6 to 2
	move 3 from 4 to 1
	move 7 from 8 to 5
	move 3 from 6 to 8
	move 11 from 2 to 1
	move 1 from 4 to 3
	move 1 from 3 to 9
	move 2 from 2 to 9
	move 8 from 5 to 4
	move 1 from 1 to 7
	move 1 from 9 to 5
	move 8 from 4 to 1
	move 1 from 6 to 8
	move 2 from 9 to 1
	move 4 from 5 to 3
	move 2 from 7 to 3
	move 40 from 1 to 2
	move 24 from 2 to 9
	move 1 from 5 to 6
	move 11 from 2 to 3
	move 9 from 3 to 5
	move 12 from 9 to 4
	move 6 from 5 to 7
	move 4 from 7 to 4
	move 2 from 5 to 1
	move 2 from 1 to 9
	move 1 from 6 to 8
	move 9 from 4 to 8
	move 6 from 4 to 9
	move 17 from 9 to 6
	move 1 from 4 to 6
	move 17 from 6 to 5
	move 1 from 1 to 4
	move 2 from 7 to 9
	move 1 from 6 to 7
	move 2 from 2 to 9
	move 2 from 7 to 2
	move 6 from 3 to 8
	move 3 from 5 to 9
	move 1 from 4 to 9
	move 2 from 3 to 7
	move 4 from 5 to 6
	move 1 from 7 to 4
	move 1 from 4 to 2
	move 1 from 7 to 5
	move 9 from 8 to 1
	move 1 from 1 to 2
	move 2 from 9 to 3
	move 7 from 2 to 7
	move 1 from 9 to 5
	move 12 from 8 to 7
	move 3 from 1 to 9
	move 2 from 6 to 4
	move 9 from 9 to 3
	move 1 from 6 to 7
	move 1 from 9 to 5
	move 1 from 6 to 1
	move 9 from 7 to 1
	move 7 from 1 to 8
	move 4 from 3 to 9
	move 5 from 7 to 1
	move 3 from 9 to 1
	move 4 from 7 to 2
	move 12 from 1 to 5
	move 2 from 9 to 4
	move 7 from 8 to 2
	move 7 from 5 to 7
	move 4 from 3 to 4
	move 1 from 8 to 1
	move 2 from 2 to 1
	move 2 from 3 to 1
	move 3 from 2 to 7
	move 13 from 5 to 4
	move 1 from 8 to 3
	move 1 from 3 to 8
	move 1 from 3 to 5
	move 1 from 8 to 7
	move 17 from 4 to 8
	move 5 from 2 to 6
	move 2 from 1 to 6
	move 5 from 6 to 3
	move 9 from 7 to 1
	move 4 from 4 to 3
	move 1 from 6 to 2
	move 4 from 7 to 4
	move 1 from 6 to 5
	move 2 from 3 to 2
	move 15 from 1 to 4
	move 6 from 5 to 4
	move 4 from 3 to 5
	move 4 from 5 to 2
	move 2 from 2 to 4
	move 11 from 8 to 1
	move 2 from 8 to 3
	move 5 from 3 to 7
	move 4 from 2 to 8
	move 2 from 2 to 9
	move 4 from 7 to 8
	move 11 from 4 to 6
	move 2 from 5 to 4
	move 3 from 6 to 9
	move 4 from 1 to 4
	move 15 from 4 to 9
	move 1 from 7 to 3
	move 2 from 1 to 2
	move 6 from 4 to 5
	move 11 from 8 to 2
	move 16 from 9 to 4
	move 2 from 9 to 1
	move 4 from 2 to 3
	move 8 from 4 to 9
	move 1 from 8 to 7
	move 5 from 4 to 7
	move 6 from 7 to 3
	move 10 from 9 to 5
	move 5 from 3 to 1
	move 1 from 1 to 4
	move 5 from 1 to 9
	move 5 from 1 to 7
	move 5 from 4 to 1
	move 4 from 1 to 6
	move 3 from 1 to 9
	move 10 from 5 to 9
	move 2 from 7 to 1
	move 5 from 3 to 6
	move 4 from 5 to 7
	move 4 from 2 to 6
	move 2 from 5 to 6
	move 5 from 2 to 7
	move 18 from 6 to 1
	move 5 from 9 to 2
	move 7 from 9 to 6
	move 16 from 1 to 7
	move 4 from 6 to 1
	move 1 from 2 to 6
	move 2 from 2 to 6
	move 1 from 2 to 4
	move 4 from 9 to 3
	move 1 from 2 to 8
	move 5 from 7 to 5
	move 2 from 9 to 3
	move 1 from 5 to 9
	move 7 from 3 to 4
	move 1 from 9 to 7
	move 8 from 1 to 9
	move 1 from 8 to 9
	move 3 from 6 to 9
	move 17 from 7 to 5
	move 3 from 4 to 8
	move 3 from 4 to 2
	move 3 from 8 to 3
	move 3 from 3 to 7
	move 7 from 9 to 3
	move 6 from 5 to 9
	move 4 from 9 to 3
	move 10 from 7 to 2
	move 15 from 5 to 2
	move 4 from 6 to 3
	move 1 from 3 to 2
	move 23 from 2 to 5
	move 2 from 4 to 6
	move 2 from 6 to 7
	move 1 from 7 to 2
	move 1 from 6 to 9
	move 5 from 9 to 8
	move 3 from 8 to 7
	move 5 from 2 to 6
	move 2 from 2 to 3
	move 2 from 6 to 3
	move 3 from 6 to 2
	move 3 from 6 to 8
	move 10 from 5 to 9
	move 2 from 7 to 5
	move 1 from 5 to 8
	move 13 from 9 to 5
	move 6 from 5 to 6
	move 1 from 6 to 1
	move 1 from 7 to 3
	move 1 from 7 to 3
	move 13 from 5 to 6
	move 3 from 3 to 5
	move 1 from 2 to 1
	move 4 from 8 to 9
	move 2 from 2 to 6
	move 2 from 5 to 3
	move 2 from 3 to 6
	move 5 from 6 to 4
	move 9 from 5 to 9
	move 10 from 6 to 9
	move 1 from 1 to 7
	move 3 from 3 to 9
	move 1 from 8 to 1
	move 3 from 6 to 3
	move 1 from 7 to 6
	move 1 from 8 to 7
	move 2 from 6 to 1
	move 2 from 6 to 4
	move 3 from 4 to 6
	move 2 from 1 to 4
	move 10 from 9 to 6
	move 6 from 4 to 9
	move 17 from 9 to 1
	move 4 from 9 to 5
	move 19 from 1 to 7
	move 4 from 5 to 6
	move 1 from 9 to 3
	move 5 from 3 to 4
	move 5 from 4 to 8
	move 17 from 6 to 9
	move 17 from 9 to 2
	move 1 from 6 to 1
	move 1 from 1 to 2
	move 1 from 8 to 3
	move 2 from 3 to 2
	move 5 from 7 to 1
	move 1 from 7 to 3
	move 5 from 2 to 9
	move 4 from 8 to 2
	move 2 from 7 to 8
	move 3 from 9 to 3
	move 7 from 3 to 9
	move 2 from 8 to 7
	move 8 from 2 to 9
	move 5 from 9 to 6
	move 4 from 3 to 9
	move 11 from 2 to 3
	move 2 from 6 to 5
	move 1 from 9 to 4
	move 10 from 7 to 3
	move 3 from 1 to 8
	move 2 from 6 to 7
	move 15 from 3 to 8
	move 2 from 3 to 2
	move 2 from 1 to 3
	move 14 from 9 to 6
	move 1 from 4 to 9
	move 14 from 6 to 3
	move 5 from 7 to 2
	move 2 from 9 to 2
	move 1 from 5 to 3
	move 1 from 5 to 8
	move 12 from 3 to 7
	move 13 from 7 to 8
	move 1 from 6 to 7
	move 5 from 2 to 6
	move 1 from 6 to 2
	move 1 from 7 to 6
	move 4 from 6 to 8
	move 31 from 8 to 7
	move 15 from 7 to 8
	move 7 from 7 to 5
	move 4 from 2 to 3
	move 1 from 6 to 2
	move 3 from 5 to 8
	move 9 from 7 to 4
	move 2 from 2 to 9
	move 4 from 5 to 6
	move 13 from 3 to 9
	move 3 from 3 to 5
	move 13 from 9 to 1
	move 1 from 3 to 2
	move 2 from 6 to 5
	move 1 from 3 to 4
	move 2 from 6 to 5
	move 1 from 9 to 1
	move 6 from 8 to 9
	move 5 from 5 to 2
	move 2 from 9 to 8
	move 2 from 1 to 6
	move 1 from 9 to 4
	move 12 from 8 to 4
	move 2 from 6 to 9
	move 11 from 4 to 3
	move 9 from 4 to 2
	move 4 from 9 to 7
	move 2 from 5 to 6
	move 8 from 3 to 4
	move 2 from 3 to 9
	move 2 from 8 to 9
	move 4 from 4 to 9
	move 2 from 6 to 7
	move 1 from 3 to 7
	move 2 from 9 to 1
	move 5 from 4 to 2
	move 9 from 1 to 8
	move 1 from 4 to 9
	move 4 from 9 to 3
	move 1 from 3 to 6
	move 4 from 8 to 7
	move 1 from 3 to 6
	move 4 from 1 to 7
	move 1 from 3 to 8
	move 1 from 1 to 8
	move 2 from 6 to 7
	move 2 from 9 to 1
	move 1 from 4 to 5
	move 1 from 1 to 5
	move 11 from 8 to 4
	move 12 from 2 to 8
	move 1 from 9 to 8
	move 2 from 4 to 5
	move 1 from 1 to 8
	move 5 from 2 to 1
	move 1 from 3 to 2
	move 9 from 7 to 3
	move 6 from 7 to 5
	move 1 from 3 to 4
	move 1 from 5 to 1
	move 4 from 2 to 5
	move 4 from 4 to 1
	move 2 from 7 to 3
	move 3 from 4 to 1
	move 6 from 3 to 7
	move 9 from 8 to 7
	move 3 from 8 to 7
	move 11 from 5 to 9
	move 2 from 4 to 8
	move 5 from 8 to 7
	move 1 from 9 to 8
	move 12 from 9 to 5
	move 1 from 4 to 5
	move 5 from 1 to 8
	move 6 from 8 to 3
	move 1 from 3 to 8
	move 3 from 7 to 9
	move 4 from 7 to 6
	move 3 from 1 to 3
	move 3 from 1 to 6
	move 1 from 8 to 1
	move 7 from 6 to 2
	move 3 from 1 to 8
	move 7 from 3 to 4
	move 3 from 4 to 1
	move 1 from 4 to 2
	move 3 from 1 to 2
	move 1 from 7 to 6
	move 1 from 8 to 5
	move 9 from 5 to 3
	move 1 from 6 to 9
	move 11 from 3 to 6
	move 1 from 4 to 1
	move 1 from 3 to 4
	move 8 from 6 to 9
	move 1 from 3 to 1
	move 1 from 9 to 1
	move 2 from 6 to 2
	move 5 from 5 to 7
	move 5 from 9 to 3
	move 2 from 8 to 5
	move 1 from 1 to 2
	move 1 from 9 to 1
	move 15 from 7 to 4
	move 1 from 1 to 6
	move 1 from 6 to 9
	move 3 from 9 to 3
	move 1 from 3 to 5
	move 5 from 5 to 3
	move 9 from 2 to 9
	move 5 from 4 to 1
	move 1 from 6 to 7
	move 7 from 9 to 3
	move 1 from 4 to 7
	move 1 from 9 to 6
	move 1 from 6 to 5
	move 2 from 1 to 4
	move 3 from 9 to 3
	move 1 from 5 to 6
	move 7 from 4 to 3
	move 1 from 9 to 3
	move 16 from 3 to 1
	move 9 from 1 to 3
	move 5 from 4 to 2
	move 1 from 6 to 9
	move 12 from 1 to 9
	move 3 from 2 to 9
	move 5 from 7 to 3
	move 2 from 4 to 8
	move 2 from 7 to 2
	move 12 from 3 to 5
	move 6 from 2 to 9
	move 12 from 3 to 1
	move 2 from 8 to 6
	move 1 from 6 to 1
	move 6 from 5 to 8
	move 5 from 3 to 2
	move 2 from 5 to 8
	move 8 from 1 to 8
	move 13 from 9 to 7
	move 4 from 7 to 5
	move 4 from 1 to 4
	move 8 from 5 to 6
	move 1 from 1 to 6
	move 4 from 7 to 3
	move 1 from 3 to 1
	move 1 from 1 to 9
	move 4 from 9 to 5
	move 3 from 3 to 7
	move 12 from 8 to 7
	move 2 from 4 to 3
	move 2 from 6 to 9
	move 4 from 8 to 2
	move 2 from 3 to 9
	move 2 from 4 to 7
	move 3 from 5 to 7
	move 2 from 9 to 7
	move 3 from 6 to 1
	move 4 from 6 to 7
	move 1 from 5 to 4
	move 1 from 9 to 3
	move 12 from 2 to 5
	move 4 from 9 to 7
	move 11 from 5 to 1
	move 1 from 6 to 5
	move 1 from 1 to 4
	move 10 from 1 to 2
	move 2 from 5 to 1
	move 1 from 3 to 5
	move 7 from 2 to 5
	move 8 from 7 to 8
	move 2 from 2 to 8
	move 3 from 9 to 4
	move 5 from 4 to 3
	move 1 from 5 to 7
	move 3 from 7 to 1
	move 3 from 5 to 8
	move 1 from 2 to 5
	move 12 from 7 to 6
	move 4 from 1 to 3
	move 2 from 5 to 6
	move 7 from 3 to 7
	move 14 from 6 to 4
	move 1 from 5 to 6
	move 3 from 1 to 3
	move 4 from 3 to 2
	move 2 from 5 to 8
	move 11 from 7 to 4
	move 7 from 4 to 5
	move 1 from 3 to 4
	move 1 from 5 to 6
	move 14 from 8 to 7
	move 11 from 7 to 3
	move 2 from 2 to 6
	move 1 from 2 to 3
	move 5 from 5 to 4
	move 4 from 6 to 4
	move 8 from 7 to 8
	move 3 from 7 to 3
	move 1 from 2 to 1
	move 5 from 8 to 2
	move 4 from 4 to 3
	move 1 from 2 to 9
	move 1 from 1 to 9
	move 3 from 2 to 1
	move 1 from 5 to 4
	move 3 from 8 to 1
	move 1 from 7 to 4
	move 4 from 3 to 9
	move 1 from 8 to 7
	move 2 from 9 to 1
	move 6 from 3 to 4
	move 28 from 4 to 7
	move 15 from 7 to 8
	move 3 from 3 to 8
	move 1 from 2 to 9
	move 2 from 3 to 2
	move 7 from 1 to 4
	move 10 from 4 to 5
	move 10 from 5 to 6
	move 3 from 8 to 2
	move 1 from 1 to 7
	move 1 from 4 to 7
	move 1 from 9 to 6
	move 9 from 6 to 7
	move 1 from 2 to 4
	move 1 from 9 to 5`

	inputData := parseCrates(input)
	commands := make([]command, 0)

	for _, line := range strings.Split(commandsText, "\n") {
		commands = append(commands, parseCommand(line))

	}
	inputData = makeComplexMoves(inputData, commands)
	printTopelements(inputData)

}

func printTopelements(input [][]string) {
	for _, column := range input {
		fmt.Printf("%s", column[len(column)-1])

	}

}

func parseCrates(input string) [][]string {
	fmt.Printf("%s\n\n", input)
	lines := strings.Split(input, "\n")
	cratesSize := (len(lines[0]) + 1) / 4
	// maxHeapSize := len(lines)
	cratesArray := make([][]string, 0)
	for i := 0; i < cratesSize; i++ {
		row := make([]string, 0)
		for j := len(lines) - 1; j >= 0; j-- {
			crate := string(lines[j][i*4+1])
			if crate != " " {
				// fmt.Println("crate = ", crate)
				row = append(row, crate)
			}

		}
		cratesArray = append(cratesArray, row)
		// fmt.Printf("%+v, %d\n", row, len(row))

	}
	// fmt.Printf("%+v\n", cratesArray)
	return cratesArray
}

type command struct {
	size int
	from int
	to   int
}

func parseCommand(line string) command {
	com := command{}
	var re = regexp.MustCompile(`\d+`)
	results := re.FindAllString(line, -1)
	com.size, _ = strconv.Atoi(results[0])
	com.from, _ = strconv.Atoi(results[1])
	com.to, _ = strconv.Atoi(results[2])
	return com

}

func makeComplexMoves(input [][]string, commands []command) [][]string {
	for _, com := range commands {
		fmt.Printf("command = %+v\n", com)
		input = makeComplexMove(input, com)
		fmt.Printf("%+v\n", input)
	}
	return input
}

// func makeMoves(input [][]string, commands []command) [][]string {
// 	for _, com := range commands {
// 		for moves := 0; moves < com.size; moves++ {
// 			fmt.Printf("command = %+v\n", com)
// 			input = makeComplexMove(input, com)
// 		}
// 		fmt.Printf("%+v\n", input)
// 	}
// 	return input
// }

// 3 A
// 2 B
// 1 C
// 0 D
func makeComplexMove(input [][]string, com command) [][]string {
	from := com.from - 1
	to := com.to - 1
	size := com.size
	fmt.Printf("from = %+v\n", input[from])

	elements := input[from][len(input[from])-size : len(input[from])]
	fmt.Printf("elements = %+v\n", elements)
	input[to] = append(input[to], elements...)
	fmt.Printf("to = %+v\n", input[to])
	input[from] = input[from][:len(input[from])-size]

	return input
}

// func makeMove(input [][]string, from, to int) [][]string {
// 	from--
// 	to--
// 	fmt.Printf("from = %+v\n", input[from])

// 	element := input[from][len(input[from])-1]
// 	fmt.Printf("element = %+v\n", element)
// 	input[to] = append(input[to], element)
// 	fmt.Printf("to = %+v\n", input[to])
// 	input[from] = input[from][:len(input[from])-1]

// 	return input
// }
