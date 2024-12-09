package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strings"
)

type position struct {
	x, y int
}

type board struct {
	field      [][]string
	maxX, maxY int
	antennas   map[string][]position
	antinodes  map[string][]position
}

func (b *board) render() {
	result := "\n"
	newfield := [][]string{}
	for y := 0; y < b.maxY; y++ {
		line := []string{}
		for x := 0; x < b.maxX; x++ {
			line = append(line, ".")

		}
		newfield = append(newfield, line)
	}
	allAntennasTypes := []string{}
	for k := range b.antennas {
		allAntennasTypes = append(allAntennasTypes, k)
	}
	sort.Strings(allAntennasTypes)
	for _, k := range allAntennasTypes {
		for _, antinode := range b.antinodes[k] {
			newfield[antinode.y][antinode.x] = "#"
		}
		for _, ant := range b.antennas[k] {
			newfield[ant.y][ant.x] = k
		}
	}

	for _, l := range newfield {
		result += strings.Join(l, "") + "\n"

	}

	fmt.Println(result)
	log.Println(b.countAntinodes())
}

func (b *board) countAntinodes() int {
	allAntinodes := []position{}
	for _, antinodes := range b.antinodes {
		for _, antinode := range antinodes {
			if !slices.Contains(allAntinodes, antinode) {
				allAntinodes = append(allAntinodes, antinode)
			}
		}
	}
	return len(allAntinodes)
}

func (b *board) findAntinodes(antId string, iterations int) {
	for i := 0; i < len(b.antennas[antId]); i++ {
		for j := 0; j < len(b.antennas[antId]); j++ {
			if i == j {
				continue
			}
			antena1 := b.antennas[antId][i]
			antena2 := b.antennas[antId][j]
			dx := antena2.x - antena1.x
			dy := antena2.y - antena1.y
			for iteration := 0; iteration <= iterations; iteration++ {
				antinode := position{antena1.x - iteration*dx, antena1.y - iteration*dy}
				if antinode.x >= 0 && antinode.x < b.maxX && antinode.y >= 0 && antinode.y < b.maxY {
					if _, ok := b.antinodes[antId]; !ok {
						b.antinodes[antId] = make([]position, 0)
					}
					b.antinodes[antId] = append(b.antinodes[antId], position{antinode.x, antinode.y})
				} else {
					break
				}
			}

		}
	}
}

func (b *board) read(input string) {
	b.antennas = make(map[string][]position)
	b.antinodes = make(map[string][]position)
	b.maxY = len(strings.Split(input, "\n")) - 1
	b.maxX = len(strings.Split(input, "\n")[0])
	for y, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		newLine := []string{}
		for x, c := range line {
			newLine = append(newLine, string(c))
			if string(c) != "." {
				if _, ok := b.antennas[string(c)]; !ok {
					b.antennas[string(c)] = make([]position, 0)
				}
				b.antennas[string(c)] = append(b.antennas[string(c)], position{x, y})
			}
		}
		b.field = append(b.field, newLine)

	}
}

func task1(input string) int {
	b := board{}
	b.read(input)
	allAntennasTypes := []string{}
	for k := range b.antennas {
		allAntennasTypes = append(allAntennasTypes, k)
		b.findAntinodes(k, 1)
	}
	b.render()

	log.Println(b.antennas)

	return 0
}

func task2(input string) int {
	b := board{}
	b.read(input)
	allAntennasTypes := []string{}
	for k := range b.antennas {
		allAntennasTypes = append(allAntennasTypes, k)
		b.findAntinodes(k, 100)
	}
	b.render()

	log.Println(b.antennas)

	return 0
}

func main() {
	// main function
	file, err := os.ReadFile("2024/07/input.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	input := string(file)
	log.Println("task2: ", task2(input))
}
