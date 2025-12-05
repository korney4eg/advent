package main

import (
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

var operations int

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	input := string(f)

	ranges, _ := parseInput(input)
	// fmt.Println("ranges", ranges)
	// fmt.Println("ingridiences", ingridiences)
	// fmt.Println("task1", countFresh(ranges, ingridiences))
	fmt.Println("task2", countFreshRanges(ranges))
	fmt.Println("total operations", operations)
}

func parseInput(input string) (ranges [][]int, ingridiences []int) {
	ingridiencesStarted := false
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			ingridiencesStarted = true
			continue
		}
		if !ingridiencesStarted {
			r := strings.Split(line, "-")
			ri0, _ := strconv.Atoi(r[0])
			ri1, _ := strconv.Atoi(r[1])
			ri := []int{ri0, ri1}
			ranges = append(ranges, ri)
		} else {
			ing, _ := strconv.Atoi(line)
			ingridiences = append(ingridiences, ing)
		}
	}
	return ranges, ingridiences
}

func countFresh(ranges [][]int, ingridiences []int) (total int) {
	for _, iningridient := range ingridiences {
		for _, r := range ranges {
			operations++
			if isIngridientInRange(iningridient, r) {
				total++
				break
			}
		}

	}
	return total
}

func isIngridientInRange(ingridient int, iRange []int) bool {
	operations++
	if ingridient >= iRange[0] && ingridient <= iRange[1] {
		return true
	}
	return false
}

func countFreshRanges(ranges [][]int) (total int) {
	joinedRanges := [][]int{}
	slices.SortFunc(ranges, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})
	for _, r := range ranges {
		added := false
		for i, ri := range joinedRanges {
			if isIngridientInRange(r[0], ri) || isIngridientInRange(r[1], ri) {
				operations++
				min := int(math.Min(float64(r[0]), float64(ri[0])))
				operations++
				max := int(math.Max(float64(r[1]), float64(ri[1])))
				operations++
				joinedRanges[i] = []int{min, max}
				added = true
			}
		}
		operations++
		if !added {
			joinedRanges = append(joinedRanges, r)
		}

	}
	for _, ri := range joinedRanges {
		total += ri[1] - ri[0] + 1
		fmt.Println(ri)
	}
	return total
}
