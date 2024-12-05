package main

import (
	"log"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func getInput(input string) (orderRules []string, pages [][]string) {
	rulesFinished := false
	for _, line := range strings.Split(input, "\n") {
		if line == "" && !rulesFinished {
			rulesFinished = true
			continue
		}
		if !rulesFinished {
			orderRules = append(orderRules, line)
		} else {
			pages = append(pages, strings.Split(line, ","))
		}
	}
	return orderRules, pages
}

func comparePages(orderRules []string, page1, page2 string) bool {
	for _, rule := range orderRules {
		more, less := strings.Split(rule, "|")[0], strings.Split(rule, "|")[1]
		if page1 == more && page2 == less {
			return true
		}
	}
	return false
}

func findMiddlePage(pagesList [][]string) int {
	sum := 0
	for i := 0; i < len(pagesList); i++ {
		middle := len(pagesList[i]) / 2
		num, _ := strconv.Atoi(pagesList[i][middle])
		sum += num
	}

	return sum
}

func arePagesOrdered(orderRules []string, pages []string) bool {
	for i := 0; i < len(pages)-1; i++ {
		if !comparePages(orderRules, string(pages[i]), string(pages[i+1])) {
			return false
		}

	}
	return true
}

func checkOrder(orderRules []string, pagesList [][]string) (orderedPages [][]string, unorderedPages [][]string) {
	for _, pages := range pagesList {
		pagesOrdered := arePagesOrdered(orderRules, pages)
		if pagesOrdered {
			orderedPages = append(orderedPages, pages)
		} else {
			unorderedPages = append(unorderedPages, pages)
		}
	}
	return orderedPages, unorderedPages
}

type pageP struct {
	id         string
	orderRules []string
}
type pagesP []*pageP

func (p pagesP) Len() int      { return len(p) }
func (p pagesP) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type ByOrder struct{ pagesP }

func (s ByOrder) Less(i, j int) bool {
	return comparePages(s.pagesP[i].orderRules, s.pagesP[i].id, s.pagesP[j].id)
}

func orderPages(orderRules []string, pagesList []string) (orderedPages []string) {
	allPages := []*pageP{}
	allRulesS := []string{}
	for _, orderRule := range orderRules {
		left, right := strings.Split(orderRule, "|")[0], strings.Split(orderRule, "|")[1]
		if !slices.Contains(allRulesS, left) {
			allRulesS = append(allRulesS, left)
		}
		if !slices.Contains(allRulesS, right) {
			allRulesS = append(allRulesS, right)
		}
	}
	for _, page := range allRulesS {
		p := &pageP{id: page, orderRules: orderRules}
		allPages = append(allPages, p)
	}
	sort.Sort(ByOrder{allPages})
	for _, page := range allPages {
		if slices.Contains(pagesList, page.id) {
			orderedPages = append(orderedPages, page.id)
		}
	}
	return orderedPages

}

func main() {
	input := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	// file, err := os.ReadFile("input.txt")
	//
	// if err != nil {
	// 	log.Fatalf("failed to open")
	// }
	//
	// input = string(file)
	orderRules, pages := getInput(input)

	log.Println("Order Rules", strings.Join(orderRules, "\n"))
	log.Println("Pages:\n", pages)
	_, unorderedPages := checkOrder(orderRules, pages)
	log.Println("unorderedPages", unorderedPages)

	orderedPages := [][]string{}
	for _, unorderedPage := range unorderedPages {
		orderedPages = append(orderedPages, orderPages(orderRules, unorderedPage))
		log.Println("original list:", unorderedPage)
		log.Println("fixed list   :", orderPages(orderRules, unorderedPage))
	}
	sum := findMiddlePage(orderedPages)
	log.Println("Sum", sum)
}
