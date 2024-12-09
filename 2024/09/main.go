package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func isDiskDifragmented(input []string) bool {
	line := strings.Join(input, "")
	re := regexp.MustCompile(`^\d+\.+$`)
	return re.MatchString(line)
}

func diskMap(input string) string {
	result := ""
	isFile := true
	i := 0
	representation := ""
	for _, c := range input {
		size, _ := strconv.Atoi(string(c))
		if isFile {
			representation = fmt.Sprintf("%d", i)
			isFile = false
			i++
		} else {
			representation = "."
			isFile = true
		}
		for j := 0; j < size; j++ {
			result += representation
		}

	}
	return result
}

func defragmentDisk(input string) string {
	diskTable := strings.Split(input, "")
	for i := len(diskTable) - 1; i >= 0; i-- {
		if isDiskDifragmented(diskTable) {
			break
		}
		if string(diskTable[i]) != "." {
			firstDot := slices.Index(diskTable, ".")
			diskTable[i], diskTable[firstDot] = diskTable[firstDot], diskTable[i]
		}
		log.Println(strings.Join(diskTable, ""))
	}
	return strings.Join(diskTable, "")

}

func getCheckSum(input string) int {
	checkSum := 0
	for i, c := range input {
		if string(c) == "." {
			break
		}
		id, _ := strconv.Atoi(string(c))
		checkSum += id * i

	}
	return checkSum
}

func task1(input string) int {
	dm := diskMap(input)
	log.Println(getCheckSum(defragmentDisk(dm)))

	return 0
}

func main() {
	file, err := os.ReadFile("2024/09/input.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	input := string(file)
	// input = `2333133121414131402`
	log.Println("task1: ", task1(input))
}
