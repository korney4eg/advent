package main

import (
    "advent/2022/12/nodes"
    "advent/utils"
    "fmt"
    "github.com/fzipp/astar"
    "image"
    "log"
    "math"
    "os"
    "strings"
)

func main() {
    maze := graph{}
    maze = append(maze, utils.ReadFileToStringsList("input.txt")...)

    //mazeMap := utils.ReadFileToStringsList("input.txt")
    start := findLetterCoordinatesInMap("S", maze)
    log.Println("Start", start)
    destinationString := "z"
    //dest := findLetterCoordinatesInMap(destinationString, maze)
    allDests := findAllLetterCoordinatesInMap(destinationString, maze)

    // Find the shortest path
    for _, point := range allDests {
        path := astar.FindPath[image.Point](maze, start, point, nodeDist, nodeDist)
        if path == nil || len(path) < 0 {
            continue
        }
        for _, p := range path {

            maze.put(p, ".", destinationString)
        }
        maze.print()
        log.Println(len(path) - 1)
        os.Exit(0)
    }

    //if path == nil {
    //    log.Println("No path found")
    //}
    // Mark the path with dots before printing

}

func findLowestScore(nodes []image.Point, dest image.Point) image.Point{
    returnNode := nodes[0]
    for _, node := range nodes {
        if nodeDist(returnNode,dest) < nodeDist(node,dest) {
            returnNode = node
        }
    }
    return returnNode
}
func reconstructPath(cameFrom map[image.Point][]image.Point, current image.Point){
total_path := []image.Point{current}

for current := range cameFrom{

}
current := cameFrom[current]
total_path.prepend(current)
return total_path
}
func aStar(start, destination image.Point) {
    openSet := []image.Point{start}
    cameFrom := make(map[image.Point][]image.Point)
    gScore := make(map[image.Point]int)
    gScore[start] = 0
    fScore := make(map[image.Point]float64)
    fScore[start] = nodeDist(start, destination)
    while len(openSet) > 0 {
        currentNode := findLowestScore(openSet, destination)
        if currentNode == destination
    }
}
func findLetterCoordinatesInMap(letter string, mazeMap []string) image.Point {
    for y, line := range mazeMap {
        for x, row := range line {
            if letter == string(row) {
                return image.Point{X: x, Y: y}
            }
        }
    }
    return image.Point{}
}
func findAllLetterCoordinatesInMap(letter string, mazeMap []string) []image.Point {
    coordinates := []image.Point{}
    for y, line := range mazeMap {
        for x, row := range line {
            if letter == string(row) {
                coordinates = append(coordinates, image.Point{X: x, Y: y})
            }
        }
    }
    return coordinates
}

// nodeDist is our cost function. We use points as nodes, so we
// calculate their Euclidean distance.
func nodeDist(p, q image.Point) float64 {
    d := q.Sub(p)
    return math.Sqrt(float64(d.X*d.X + d.Y*d.Y))
}

type graph []string

// Neighbours implements the astar.Graph[Node] interface (with Node = image.Point).
func (g graph) Neighbours(p image.Point) []image.Point {
    offsets := []image.Point{
        image.Pt(0, -1), // North
        image.Pt(1, 0),  // East
        image.Pt(0, 1),  // South
        image.Pt(-1, 0), // West
    }
    res := make([]image.Point, 0, 4)
    for _, off := range offsets {
        q := p.Add(off)
        if g.hasConnections(p, q) {
            res = append(res, q)
        }
    }
    return res
}

func (g graph) hasConnections(p, q image.Point) bool {
    if q.Y < 0 || q.X < 0 || q.Y >= len(g) || q.X >= len(g[q.Y]) {
        return false
    }
    if g[p.Y][p.X] == g[q.Y][q.X] {
        return true
    }
    pHeight := nodes.GetNodeHeight(string(g[p.Y][p.X]))
    qHeight := nodes.GetNodeHeight(string(g[q.Y][q.X]))
    //log.Printf("p(%s) %d: %v", string(g[p.Y][p.X]), pHeight, p)
    //log.Printf("q(%s) %d: %v", string(g[q.Y][q.X]), qHeight, q)
    //log.Printf("Do they have connection? %v", pHeight-qHeight <= 1)
    return math.Abs(float64(pHeight-qHeight)) <= 1
}

func replaceLineWithDots(line, destination string) string {
    outputLine := ""
    for _, letter := range line {
        if nodes.GetNodeHeight(string(letter)) <= nodes.GetNodeHeight(destination) || string(letter) != "v" {
            outputLine += "."
        } else {
            outputLine += string(letter)
        }
    }
    return outputLine
}

func (g graph) put(p image.Point, c string, destination string) {

    g[p.Y] = g[p.Y][:p.X] + strings.ToUpper(string(g[p.Y][p.X])) + g[p.Y][p.X+1:]
}

func (g graph) print() {
    for _, row := range g {
        fmt.Println(row)
    }
}
