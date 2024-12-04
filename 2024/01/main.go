package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getInput(input string) (leftList, rightList []int) {
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		re := regexp.MustCompile(`^([0-9]*)\s*([0-9]*)`)
		line = re.ReplaceAllString(line, `$1 $2`)
		ls := strings.Split(line, " ")[0]
		rs := strings.Split(line, " ")[1]
		ln, _ := strconv.Atoi(ls)
		leftList = append(leftList, ln)
		rn, _ := strconv.Atoi(rs)
		rightList = append(rightList, rn)
	}
	return leftList, rightList
}

func getDistances(leftList, rightList []int) (distances []int) {
	for i := 0; i < len(leftList); i++ {
		distance := int(math.Abs(float64(rightList[i] - leftList[i])))
		distances = append(distances, distance)
	}
	return distances
}

func getSimilarities(leftList, rightList []int) (similarities []int) {
	simMap := make(map[int]int)
	for i := 0; i < len(leftList); i++ {
		log.Printf("leftList[i]=%d, simMap=%+v", leftList[i], simMap)
		val, ok := simMap[leftList[i]]
		if !ok {
			founds := 0
			for j := 0; j < len(rightList); j++ {
				if leftList[i] == rightList[j] {
					founds++
				}
			}
			simMap[leftList[i]] = founds * leftList[i]
			log.Printf("set simMap=%+v", simMap)
			val = simMap[leftList[i]]
		}
		similarities = append(similarities, val)
		fmt.Println("similarities", similarities)
	}
	return similarities
}

func main() {

	file, err := os.ReadFile("test.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	input := string(file)
	// 	input := `3   4
	// 4   3
	// 2   5
	// 1   3
	// 3   9
	// 3   3`
	leftList, rightList := getInput(input)
	distances := getSimilarities(leftList, rightList)
	sum := 0
	for _, distance := range distances {
		sum += distance
	}
	fmt.Println(sum)
}
