package main

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

var operations int
var cache = make(map[int][]int)
var hits, misses int

func proceed(num int) (result []int) {
	operations++
	if num == 0 {
		return []int{1}
	} else if len(strconv.Itoa(num))%2 == 0 {
		operations++
		numStr := strconv.Itoa(num)
		operations++
		first := numStr[:len(numStr)/2]
		operations++
		f, _ := strconv.Atoi(first)
		operations++
		second := numStr[len(numStr)/2:]
		operations++
		s, _ := strconv.Atoi(second)
		return []int{f, s}
	}
	operations++
	newNum := num * 2024
	operations++

	return []int{newNum}

}

func blink(input []int) (result []int) {
	for _, num := range input {
		newList := []int{}
		hits++
		res, ok := cache[num]
		if !ok {
			newList = proceed(num)
			cache[num] = newList
			misses++
		} else {
			newList = res
		}

		result = append(result, newList...)

	}
	// slices.Sort(result)
	return result
}

func task1(input string) int {
	stonesList := strings.Split(input, " ")
	stones := []int{}
	for _, stone := range stonesList {
		operations++
		s, _ := strconv.Atoi(stone)
		stones = append(stones, s)
	}
	nums := 25
	for i := 0; i < nums; i++ {
		stones = blink(stones)
		// log.Println(stones)
		log.Println(i, "/", nums)
		log.Println("miss rate", misses, "/", hits, "=", float64(misses)/float64(hits))
		log.Println("operations", operations)

	}
	return len(stones)
}

func main() {
	input := "125 17"
	input = "20 82084 1650 3 346355 363 7975858 0"
	task1Result := task1(input)
	log.Println(task1Result)
}
