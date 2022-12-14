package main

import (
	"advent/2022/13/signals"
	"advent/utils"
	"log"
)

func main() {
	index := 0
	pair := 0
	lines := utils.ReadFileToStringsList("2022/13/input.txt")
	for i := 0; i < len(lines); i += 3 {
		pair++
		element1 := lines[i]
		element2 := lines[i+1]
		if signals.IsRightOrder(element1, element2) {
			log.Println()
			log.Println()
			index += pair
		}
	}
	log.Println(index)

}
