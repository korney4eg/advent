package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"slices"
)

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	input := string(f)
	parsedInput := parseInput(input)

	sum := 0
	for _, inp := range parsedInput {
		ids := listAllInvalidIds(inp, 2)
		fmt.Printf("for %s got: %v\n", inp, ids)
		for _, id := range ids {
			sum += id
		}
	}

	fmt.Println("First", sum)
	fmt.Println("============================")

	sum = 0
	for _, inp := range parsedInput {
		ids := getAllInvalidIds(inp)
		fmt.Printf("for %s got: %v\n", inp, ids)
		for _, id := range ids {
			sum += id
		}
	}

	fmt.Println("First", sum)
}

func parseInput(input string) []string {
	joined := strings.ReplaceAll(input, "\n", "")
	return strings.Split(joined, ",")
}

func getAllInvalidIds(input string) []int {
	out := []int{}
	second := strings.Split(input, "-")[1]
	// fmt.Println(second)
	for i := 2; i <= len(second); i++ {
		fmt.Printf("%s  i: %d\n", input, i)
		ids := listAllInvalidIds(input, i)
		fmt.Printf("for %s got: %v, i: %d\n", input, ids, i)
		for _, id := range ids {
			if !slices.Contains(out, id) {
				out = append(out, id)
			}
		}
	}
	return out
}

func listAllInvalidIds(input string, n int) []int {
	invalidIdsCounter := 0
	ids := []int{}
	firstN, _ := strconv.Atoi(strings.Split(input, "-")[0])
	nextInvalidId := strings.Split(input, "-")[0]
	finish := strings.Split(input, "-")[1]
	if !isIdValid(nextInvalidId, n) {
		// fmt.Println("Is not valid from first try", nextInvalidId, n)
		invalidIdsCounter++
		n, _ := strconv.Atoi(nextInvalidId)
		ids = append(ids, n)
	}
	for {
		nextInvalidId = getNextInvalidId(nextInvalidId, n)
		// fmt.Printf("nexId '%s', finish: '%s'\n", nextInvalidId, finish)
		nextN, _ := strconv.Atoi(nextInvalidId)
		finishN, _ := strconv.Atoi(finish)
		if nextN > finishN || nextN < firstN {
			break
		}
		ids = append(ids, nextN)
	}
	return ids
}

func getNextInvalidId(id string, n int) string {
	num, _ := strconv.Atoi(id)
	if !isIdValid(id, n) {
		num++
		id = strconv.Itoa(num)
	}
	// fmt.Println("looking for next in id", id)
	if len(id)%n != 0 {
		// find next 10^n
		nextNum10 := int(math.Pow(10, float64(len(id))))
		return getNextInvalidId(strconv.Itoa(nextNum10), n)
	}
	min := "9999999"
	max := ""
	// fmt.Println("id:", id, "n:", n)
	// second := id[len(id)/n:]
	// if first > second {
	// 	return first + first
	// }
	part := ""
	length := len(id) / n
	first := id[0 : len(id)/n]
	for i := 0; i < n; i++ {
		start := i * length
		end := start + length
		fmt.Println("mmin", min, "element", id[start:end])
		if min > id[start:end] {
			min = id[start:end]
		}
		if max < id[start:end] {
			max = id[start:end]
		}
		fmt.Println("mmmmin", min, "max", max)
	}
	if first == max {
		return strings.Repeat(first, n)
	}
	firstN, _ := strconv.Atoi(first)
	// secondN, _ := strconv.Atoi(second)
	part = strconv.Itoa(firstN + 1)
	fmt.Println("min", min, "id", id, "next", strings.Repeat(part, n))

	return strings.Repeat(part, n)
}

func isIdValid(id string, n int) bool {
	for i := 2; i <= n; i++ {
		if len(id)%i != 0 {
			continue
		}
		if strings.Repeat(id[0:len(id)/i], i) == id {
			return false
		}
	}
	return true
}
