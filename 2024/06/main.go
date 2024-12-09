package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/fatih/color"
	"github.com/veandco/go-sdl2/sdl"
)

var MaxScreenWidth = 800
var MaxScreenHeight = 800
var tyleSize int

type position struct {
	x, y int
}

type direction int

func (d direction) String() string {
	switch d {
	case up:
		return "^"
	case down:
		return "⌄"
	case left:
		return "<"
	case right:
		return ">"
	}
	return "unknown"
}

func (g *guard) getNextPosition() position {
	p := g.position
	switch g.direction {
	case up:
		return position{x: p.x, y: p.y - 1}
	case down:
		return position{x: p.x, y: p.y + 1}
	case left:
		return position{x: p.x - 1, y: p.y}
	case right:
		return position{x: p.x + 1, y: p.y}
	}
	return p
}
func rotate(d direction) direction {
	directions := []direction{up, right, down, left}
	directionIndex := slices.Index(directions, d)
	if directionIndex+1 == len(directions) {
		d = directions[0]
		return d
	}
	d = directions[directionIndex+1]
	return d
}

const (
	up direction = iota
	right
	down
	left
)

type guard struct {
	position  position
	direction direction
}

type board struct {
	guard          guard
	wallsPosition  []position
	maxRow, maxCol int
	visited        map[position]int
	visitedList    []position
	field          [][]string
	recourced      bool
}

func (b *board) addVisited(pos position) {
	if _, ok := b.visited[pos]; !ok {
		b.visited[pos] = 1
		b.visitedList = append(b.visitedList, pos)
	} else {
		b.visited[pos] += 1
	}
}

func (b *board) findPlusesOnDir(pos position, dx, dy int) bool {
	maxX, maxY := 0, 0
	if dy > 0 {
		maxY = b.maxRow - 1
	} else if dy < 0 {
		maxY = -1
	}
	if dx > 0 {
		maxX = b.maxCol - 1
	} else if dy < 0 {
		maxX = -1
	}
	if dx == 0 && dy != 0 {
		for y := pos.y; y != maxY; y += dy {

			if b.field[y][pos.x] == "#" {
				return true
			}
		}
	} else if dx != 0 && dy == 0 {
		for x := pos.x; x != maxX; x += dx {

			if b.field[pos.y][x] == "#" {
				return true
			}
		}

	}
	return false

}

func (b *board) read(input string) {
	lines := strings.Split(input, "\n")
	b.visited = make(map[position]int)
	b.maxRow = len(lines)
	b.maxCol = len(lines[0])

	for i, line := range lines {
		for j, c := range line {
			if string(c) == "#" {
				wallPosition := position{x: j, y: i}
				b.wallsPosition = append(b.wallsPosition, wallPosition)
			} else if string(c) == "^" {
				b.guard = guard{position: position{x: j, y: i}, direction: up}
			}
		}
		b.field = append(b.field, strings.Split(strings.ReplaceAll(line, ".", " "), ""))
	}
}

func (b *board) moveGuard(rotated bool) {
	pathSymbol := "-"
	if b.guard.direction == up || b.guard.direction == down {
		pathSymbol = "|"
	}
	if rotated {
		pathSymbol = "+"
	}
	b.field[b.guard.position.y][b.guard.position.x] = pathSymbol
	b.addVisited(b.guard.position)
	b.guard.position = b.guard.getNextPosition()
}

func (b *board) isNextMoveInBoard() bool {
	nextPosition := b.guard.getNextPosition()
	if nextPosition.x < 0 || nextPosition.x >= b.maxCol || nextPosition.y < 0 || nextPosition.y >= b.maxRow {
		return false
	}
	return true
}

func (b *board) render() {
	lines := []string{}
	for i := 0; i < b.maxRow; i++ {
		lines = append(lines, strings.Join(b.field[i], ""))
	}
	fmt.Println(strings.Join(lines, "\n"))
	fmt.Println("______________________")

}
func (b *board) isNextStepWall() bool {
	nextPostion := b.guard.getNextPosition()
	for _, wall := range b.wallsPosition {
		if wall.x == nextPostion.x && wall.y == nextPostion.y {
			return true
		}
	}
	return false
}

func (b *board) update() {
	rotated := false
	if b.isNextStepWall() {
		b.guard.direction = rotate(b.guard.direction)
		rotated = true
	}

	b.moveGuard(rotated)
}
func newBoard(input string) *board {
	b := &board{}
	b.read(input)
	for true {
		b.update()
		if !b.isNextMoveInBoard() && !b.isNextStepWall() {
			b.addVisited(b.guard.position)
			break
		}
	}
	return b
}

func drawRect(renderer *sdl.Renderer, x, y int32, colour *sdl.Color) {
	rect := sdl.Rect{x, y, int32(tyleSize), int32(tyleSize)}
	renderer.SetDrawColor(colour.R, colour.G, colour.B, colour.A)
	renderer.FillRect(&rect)
}

func (b *board) renderSDL(renderer *sdl.Renderer) {
	rect := sdl.Rect{0, 0, int32(b.maxCol * tyleSize), int32(b.maxRow * tyleSize)}
	renderer.SetDrawColor(0, 0, 222, 0)
	renderer.DrawRect(&rect)
	for _, v := range b.visitedList {
		x := v.x
		y := v.y
		if b.recourced {
			drawRect(renderer, int32(x*tyleSize), int32(y*tyleSize), &sdl.Color{R: 255, G: 0, B: 0, A: 0})
		} else {
			drawRect(renderer, int32(x*tyleSize), int32(y*tyleSize), &sdl.Color{R: 0, G: 255, B: 0, A: 0})
		}
	}
	for _, wall := range b.wallsPosition {
		x := wall.x
		y := wall.y
		drawRect(renderer, int32(x*tyleSize), int32(y*tyleSize), &sdl.Color{R: 0, G: 0, B: 0, A: 0})
	}

}

func (b *board) findRecources() {
	bumpedWalls := []position{}
	for true {
		if !b.isNextMoveInBoard() {
			b.addVisited(b.guard.position)
			return
		}
		nextPosition := b.guard.getNextPosition()
		rotated := false
		if b.isNextStepWall() {
			bumpedWalls = append(bumpedWalls, nextPosition)
			b.guard.direction = rotate(b.guard.direction)
			b.addVisited(b.guard.position)
			rotated = true
		} else {
			b.addVisited(nextPosition)
			log.Println("recurrented")
			b.render()
			b.recourced = true
			return
		}
		b.moveGuard(rotated)
	}

}

func task2(input string) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(MaxScreenHeight), int32(MaxScreenWidth), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		panic(err)
	}
	defer renderer.Destroy()

	initialBoard := newBoard(input)
	tyleSize = MaxScreenWidth / initialBoard.maxCol
	running := true
	i := 1
	b := &board{}
	b.read(input)
	b.wallsPosition = append(b.wallsPosition, initialBoard.visitedList[i])
	b.field[initialBoard.visitedList[i].y][initialBoard.visitedList[i].x] = color.RedString("#")
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			if true {
			}
			switch t := event.(type) {
			case *sdl.QuitEvent: // NOTE: Please use `*sdl.QuitEvent` for `v0.4.x` (current version).
				println("Quit")
				running = false
				break
			case *sdl.KeyboardEvent:
				keyCode := t.Keysym.Sym

				if t.State == sdl.RELEASED && keyCode == sdl.K_SPACE {
					i++
					b = &board{}
					b.read(input)
					b.wallsPosition = append(b.wallsPosition, initialBoard.visitedList[i])
					b.field[initialBoard.visitedList[i].y][initialBoard.visitedList[i].x] = color.RedString("#")
					b.findRecources()
					b.render()
				}
				if t.State == sdl.RELEASED && keyCode == sdl.K_ESCAPE {
					println("Quit")
					running = false
					break
				}

			}
		}

		renderer.SetDrawColor(236, 236, 236, 0)
		renderer.Clear()
		b.renderSDL(renderer)
		renderer.Present()
		sdl.Delay(16)
		renderer.Clear()
		if i == len(initialBoard.visited)-1 {
			break
		}
	}

}

func main() {
	file, err := os.ReadFile("2024/06/test.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	input := string(file)
	task2(input)
}
