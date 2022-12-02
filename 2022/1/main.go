package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	candidate := 0
	foods := make([]int,0)

	for scanner.Scan() {
		line := scanner.Text()
		if  line == "" {
				foods = append(foods, candidate)
				candidate = 0
				continue
		}
		number, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		candidate += number
	}
	
	sort.Sort(sort.Reverse(sort.IntSlice(foods)))
	log.Printf("%d\n", foods[0]+foods[1]+foods[2])
}
