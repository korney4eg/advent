package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var operations = 0
var freeSpaceSymbol = "_"

var files []int
var freeSpaces []int
var filesF [][]string
var freeSpacesF [][]string
var fLocations map[string]int

func isDiskDifragmented(input []string) bool {
	line := strings.Join(input, "")
	re := regexp.MustCompile(`^\d+\.+$`)
	operations++
	return re.MatchString(line)
}

func diskMap(input string) (result []string) {
	fLocations = make(map[string]int)

	isFile := true
	i := 0
	for _, c := range input {
		operations++
		size, _ := strconv.Atoi(string(c))
		if isFile {
			ff := []string{}
			for j := 0; j < size; j++ {
				fLocations[fmt.Sprintf("%d", i)] = len(result)
				result = append(result, fmt.Sprintf("%d", i))
				ff = append(ff, fmt.Sprintf("%d", i))

			}
			filesF = append(filesF, ff)
			isFile = false
			i++
		} else {
			result = append(result, strings.Split(strings.Repeat(freeSpaceSymbol, size), "")...)
			freeSpacesF = append(freeSpacesF, strings.Split(strings.Repeat(freeSpaceSymbol, size), ""))
			isFile = true
		}

	}
	for j := 0; j < len(result); j++ {
		if result[j] == freeSpaceSymbol {
			freeSpaces = append(freeSpaces, j)
		} else {

			files = append(files, j)
		}
	}
	log.Println("maxId ", i)
	return result
}

func swapSpaces(input []string) (output []string) {
	a := (make([]string, len(input)))
	i := files[len(files)-1]
	j := freeSpaces[0]
	fileBlock := input[i]
	freeSpaceBlock := input[j]

	output = input[:j]
	for b, c := range input {
		operations++
		if b == i {
			a[b] = freeSpaceBlock
		} else if b == j {
			a[b] = fileBlock
		} else {
			a[b] = string(c)
		}
	}
	files = append([]int{j}, files[:len(files)-1]...)
	freeSpaces = append(freeSpaces[1:], i)

	return a[:]
}
func defragmentDisk(input []string) []string {
	diskTable := input
	// i := 0
	spaceToFree := slices.Clone(input)
	spaceToFree = spaceToFree[len(files):]
	sort.Sort(sort.Reverse(sort.StringSlice(spaceToFree)))
	spaceToFreeFixed := []string{}
	for _, c := range spaceToFree {
		symb := strings.TrimLeft(c, "0")
		if symb == "" {
			symb = "0"
		}

		spaceToFreeFixed = append(spaceToFreeFixed, symb)
	}
	inputFixed := []string{}
	for _, c := range input {
		symb := strings.TrimLeft(c, "0")
		if symb == "" {
			symb = "0"
		}

		inputFixed = append(inputFixed, symb)
	}
	log.Println("spaceToFree", spaceToFreeFixed)
	a := inputFixed[:len(files)]
	j := 0
	for i := 0; i < len(files); i++ {
		freeSymbolIndex := slices.Index(a, freeSpaceSymbol)
		if freeSymbolIndex == -1 {
			continue
		}
		a[freeSymbolIndex] = spaceToFreeFixed[j]
		j++
		// log.Println("a", a)

	}
	diskTable = a[:]
	log.Println("diskTable", diskTable)

	// for true {
	// 	operations++
	// 	if isDiskDifragmented(diskTable) {
	// 		operations++
	// 		break
	// 	}
	//
	// 	diskTable = swapSpaces(diskTable)
	// 	// log.Println(strings.Join(diskTable, ""))
	// 	i++
	// }
	return diskTable

}

func defragmentDiskFullFile(input []string) []string {
	diskTable := input

	a := make([]string, len(input))
	copy(a, input)
	location58 := slices.Index(input, "58")
	log.Println("location58", location58)
	for f := len(filesF) - 1; f > 0; f-- {
		myFiles := filesF[f]
		c := fmt.Sprintf("%d", f)
		fileIdIndex := slices.Index(input, c)
		// found := false
		// if fileIdIndex == -1 {
		// 	fileIdIndex = fLocations[c]
		// }
		// if f < 100 {
		// 	log.Println("c", c, "myFiles", myFiles)
		// }

		if fileIdIndex == -1 {
			log.Println("coulnd't find symbol", myFiles[0])
		}
		fileLen := len(myFiles)
		line := strings.Join(a, "")
		fileNewIndex := strings.Index(line, strings.Repeat(freeSpaceSymbol, fileLen))
		if fileIdIndex < fileNewIndex {
			// log.Println("fileIdIndex", fileIdIndex, "fileNewIndex", fileNewIndex, "symbol", myFiles[0])
			continue
		}
		if fileNewIndex != -1 {
			for i := 0; i < fileLen; i++ {
				a[fileNewIndex+i] = myFiles[0]
				a[fileIdIndex+i] = freeSpaceSymbol
			}

		}
		// log.Println("a", a)
	}
	// log.Println("locations", fLocations["58"])
	diskTable = a[:]
	return diskTable

}

func getCheckSum(input []string) int {
	checkSum := 0
	for i, c := range input {
		operations++
		if string(c) == freeSpaceSymbol {
			operations++
			continue
		}
		id, _ := strconv.Atoi(c)
		checkSum += id * i
		// log.Println("id", id, "i", i, "c:", c, "checkSum", checkSum)

	}
	return checkSum
}

func task2(input string) int {
	dm := diskMap(input)
	log.Println(dm)

	return getCheckSum(defragmentDiskFullFile(dm))
}

func task1(input string) int {
	dm := diskMap(input)
	log.Println(dm)

	return getCheckSum(defragmentDisk(dm))
}

func main() {
	file, err := os.ReadFile("2024/09/input2.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	input := string(file)
	// input = `2333133121414131402`
	log.Println("task2: ", task2(input))
	log.Println("operations: ", operations)
}
