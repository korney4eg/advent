package main

import (
	"advent/2022/14/maze"
	"advent/utils"
	"log"
)

func main() {
	lines := utils.ReadFileToStringsList("2022/14/input.txt")
	newMaze := maze.NewMaze(lines)
	sandUnits := 9999999
	for sandUnit := 0; sandUnit < sandUnits; sandUnit++ {
		sand := &maze.Sand{X: maze.PouringX, Y: maze.PouringY}
		for {
			//log.Println("\n" + sand.Draw(newMaze))
			//fmt.Scanln()
			if !sand.CanFallOneTile(newMaze) || sand.OutOfScreen {
				break
			}

		}
		if sand.OutOfScreen {
			log.Println(sandUnit)
			break
		}
	}
}
