package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var dict = map[string]int{"r": rockPoints, "p": paperPoints, "s": scisorsPoints}
var enemyMap = map[string]string{"A": rockSign, "B": paperSign, "C": scisorsSign}
var ourMap = []map[string]string{
	map[string]string{"X": rockSign, "Y": paperSign, "Z": scisorsSign},
	map[string]string{"X": rockSign, "Z": paperSign, "Y": scisorsSign},
	map[string]string{"Z": rockSign, "Y": paperSign, "X": scisorsSign},
	map[string]string{"Z": rockSign, "X": paperSign, "Y": scisorsSign},
	map[string]string{"Y": rockSign, "X": paperSign, "Z": scisorsSign},
	map[string]string{"Y": rockSign, "Z": paperSign, "X": scisorsSign},
}

var winMap = map[string]string{"X": string(losePoints), "Y": string(drawPoints), "Z": string(winPoints)}

const (
	rockSign    = "r"
	paperSign   = "p"
	scisorsSign = "s"

	rockPoints    = 1
	paperPoints   = 2
	scisorsPoints = 3

	drawPoints = 3
	winPoints  = 6
	losePoints = 0
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	linesList := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		linesList = append(linesList, line)
	}

	total := 0
	for _, line := range linesList {
		enemyChoice := strings.Split(line, " ")[0]
		ourWinChoice := strings.Split(line, " ")[1]
		ourChoice := GetOurChose(enemyChoice, ourWinChoice)
		sum := GetRoundPoints(enemyMap[enemyChoice], ourChoice)
		total += sum
	}
	fmt.Printf("total = %d\n", total)

}

func GetRoundPoints(enemySign, ourSign string) int {
	return dict[ourSign] + GetVictoryPoints(enemySign, ourSign)
}

func GetVictoryPoints(enemySign, ourSign string) int {
	if enemySign == ourSign {
		return drawPoints
	}
	switch enemySign + ourSign {
	case "rp":
		return winPoints
	case "rs":
		return losePoints
	case "ps":
		return winPoints
	case "pr":
		return losePoints
	case "sr":
		return winPoints
	case "sp":
		return losePoints
	}
	return 0

}
func GetOurChose(enemySign string, ourPoints string) string {
	if ourPoints == "Y" {
		return enemyMap[enemySign]
	}
	switch enemySign + ourPoints {
	case "AZ":
		return paperSign
	case "AX":
		return scisorsSign
	case "BZ":
		return scisorsSign
	case "BX":
		return rockSign
	case "CZ":
		return rockSign
	case "CX":
		return paperSign
	}
	return ""
}
