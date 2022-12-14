package maze

import (
	"log"
	"math"
	"strconv"
	"strings"
)

const PouringX = 500
const PouringY = 0

type Maze struct {
	Field [][]string
	MinX  int
	MinY  int
	MaxX  int
	MaxY  int
}

type Sand struct {
	X           int
	Y           int
	OutOfScreen bool
}

func (s *Sand) Draw(maze *Maze) string {
	currentTile := maze.Field[s.Y][s.X]
	maze.Field[s.Y][s.X] = "o"
	output := ""
	for y := maze.MinY; y <= maze.MaxY; y++ {
		output += strings.Join(maze.Field[y][maze.MinX:maze.MaxX], "") + "\n"

	}
	maze.Field[s.Y][s.X] = currentTile
	return output
}
func (s *Sand) CanFallOneTile(maze *Maze) (canContinue bool) {
	maze.MinX = int(math.Min(float64(maze.MinX), float64(s.X-1)))
	maze.MaxX = int(math.Max(float64(maze.MaxX), float64(s.X+1)))

	if maze.Field[s.Y+1][s.X] == "." {
		s.Y = s.Y + 1
		return true
	}
	if s.X > 0 && maze.Field[s.Y+1][s.X-1] == "." {
		s.Y = s.Y + 1
		s.X = s.X - 1
		return true
	}
	if s.X < len(maze.Field[0]) && maze.Field[s.Y+1][s.X+1] == "." {
		s.Y = s.Y + 1
		s.X = s.X + 1
		return true
	}

	maze.Field[s.Y][s.X] = "o"
	return false

}

func (m *Maze) addWallVector(beginVectorX, beginVectorY, endVectorX, endVectorY int) {
	minX := int(math.Min(float64(beginVectorX), float64(endVectorX)))
	minY := int(math.Min(float64(beginVectorY), float64(endVectorY)))
	maxX := int(math.Max(float64(beginVectorX), float64(endVectorX)))
	maxY := int(math.Max(float64(beginVectorY), float64(endVectorY)))
	m.MinX = int(math.Min(float64(m.MinX), float64(minX)))
	m.MinY = int(math.Min(float64(m.MinY), float64(minY)))
	m.MaxX = int(math.Max(float64(m.MaxX), float64(maxX)))
	m.MaxY = int(math.Max(float64(m.MaxY), float64(maxY)))

	if beginVectorX == endVectorX {
		for i := minY; i <= maxY; i++ {
			m.Field[i][beginVectorX] = "#"
		}
	} else if beginVectorY == endVectorY {
		for i := minX; i <= maxX; i++ {
			m.Field[beginVectorY][i] = "#"
		}
	}
}

func (m *Maze) Draw() (output string) {
	for y := m.MinY; y <= m.MaxY; y++ {
		output += strings.Join(m.Field[y][m.MinX:m.MaxX], "") + "\n"

	}
	log.Println(m.MinX, m.MaxX, m.MinY, m.MaxY)
	return output
}

func (m *Maze) getMinMaxValues(lines []string) {
	for _, line := range lines {
		for _, nums := range strings.Split(line, " -> ") {
			x, _ := strconv.Atoi(strings.Split(nums, ",")[0])
			y, _ := strconv.Atoi(strings.Split(nums, ",")[1])
			m.MinX = int(math.Min(float64(m.MinX), float64(x)))
			m.MinY = int(math.Min(float64(m.MinY), float64(y)))
			m.MaxX = int(math.Max(float64(m.MaxX), float64(x)))
			m.MaxY = int(math.Max(float64(m.MaxY), float64(y)))
		}
	}
	m.MinX = int(math.Min(float64(m.MinX), float64(PouringX)))
	m.MinY = int(math.Min(float64(m.MinY), float64(PouringY)))
	m.MaxX = int(math.Max(float64(m.MaxX), float64(PouringX)))
	m.MaxY = int(math.Max(float64(m.MaxY), float64(PouringY)))
	m.MaxY += 2
}

func (m *Maze) initiateField() {
	mazeLines := [][]string{}

	for y := 0; y <= m.MaxY; y++ {
		mazeLine := []string{}
		element := "."
		if y == m.MaxY {
			element = "#"
		}
		for x := 0; x <= 999999; x++ {
			mazeLine = append(mazeLine, element)
		}
		mazeLines = append(mazeLines, mazeLine)
	}
	mazeLines[PouringY][PouringX] = "+"
	m.Field = mazeLines
}

func NewMaze(lines []string) *Maze {
	newMaze := &Maze{}
	newMaze.MaxY = 0
	newMaze.MaxY = 0
	newMaze.MinX = 9999999
	newMaze.MinY = 9999999
	newMaze.getMinMaxValues(lines)
	newMaze.initiateField()
	for _, line := range lines {
		wallVectors := strings.Split(line, " -> ")
		for i := 0; i < len(wallVectors)-1; i++ {
			beginVectorX, _ := strconv.Atoi(strings.Split(wallVectors[i], ",")[0])
			beginVectorY, _ := strconv.Atoi(strings.Split(wallVectors[i], ",")[1])
			endVectorX, _ := strconv.Atoi(strings.Split(wallVectors[i+1], ",")[0])
			endVectorY, _ := strconv.Atoi(strings.Split(wallVectors[i+1], ",")[1])
			newMaze.addWallVector(
				beginVectorX,
				beginVectorY,
				endVectorX,
				endVectorY,
			)
		}
	}
	return newMaze
}
