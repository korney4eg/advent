package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

	sum := 0

	/*	for _, line := range linesList {
			leftPart := line[0 : len(line)/2]
			rightPart := line[len(line)/2 : len(line)]
			commonLetter, err := getCommonLetter(leftPart, rightPart)
			if err != nil {
				log.Fatal(err)
			}
			//log.Println(commonLetter, getPoints(commonLetter))
			sum += getPoints(commonLetter)

			//log.Printf("\nline %s,\n leftPart %s,\nrightPart %s,\n", line, leftPart, rightPart)
		}
		log.Println(sum)
	*/
	for i := 0; i < len(linesList)/3; i++ {
		firstPart := linesList[i*3]
		secondPart := linesList[i*3+1]
		thirdPart := linesList[i*3+2]

		commonLetter, err := getCommonLetter(firstPart, secondPart, thirdPart)

		if err != nil {
			log.Fatal(err)
		}

		sum += getPoints(commonLetter)

	}

	log.Println(sum)
}

func getCommonLetter(first, second, third string) (string, error) {
	for _, firstItem := range first {
		for _, secondItem := range second {
			for _, thirdItem := range third {
				if firstItem == secondItem && secondItem == thirdItem {
					return string(firstItem), nil
				}
			}
		}
	}

	return "", fmt.Errorf("Can not find item")
}

func getPoints(letter string) int {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(alphabet, letter) + 1
}
