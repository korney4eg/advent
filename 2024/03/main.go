package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var operations int

type word struct {
	lenght int
	color  string
}

func (s *word) colorize() string {
	return ""
}

type renderer struct {
	line       string
	highlights map[int]word
}

func (r *renderer) render(muls []string) {
	dos := regexp.MustCompile(`do\(\)`)
	line := dos.ReplaceAllString(r.line, color.GreenString("do()"))
	donts := regexp.MustCompile(`don't\(\)`)
	line = donts.ReplaceAllString(line, color.RedString("don't()"))
	for _, m := range muls {
		mr := strings.Replace(m, "(", "\\(", -1)
		mr = strings.Replace(mr, ")", "\\)", -1)
		mul := regexp.MustCompile(mr)
		line = mul.ReplaceAllString(line, color.YellowString(m))
	}
	fmt.Println(line)
}

func findCorrectInstructions(input string) (instructions []string) {
	re := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	operations++
	return re.FindAllString(input, -1)
}

func findStrs(input, word string) (dos []int) {
	re := regexp.MustCompile(word)
	for _, found := range re.FindAllStringIndex(input, -1) {
		dos = append(dos, found[0])
	}
	return dos
}

func isEnabled(index int, dos, donts []int) bool {
	enabled := true
	diff := 99999
	for _, d := range donts {
		if d > index {
			break
		}
		if index-d < diff && index-d >= 0 {
			diff = index - d
			enabled = false
		}
	}
	for _, d := range dos {
		if d > index {
			break
		}
		if index-d < diff && index-d >= 0 {
			diff = index - d
			enabled = true
		}
	}
	return enabled
}

func findEnabledInstructions(input string) (instructions []string) {
	re := regexp.MustCompile(`do\(\)|mul\([0-9]+,[0-9]+\)|don't\(\)`)
	foundElements := re.FindAllString(input, -1)
	flag := true
	for _, element := range foundElements {
		if element == "do()" {
			flag = true
		} else if element == "don't()" {
			flag = false
		} else {
			if flag {
				instructions = append(instructions, element)
			}
		}
	}

	return instructions
}

func mul(input string) int {
	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	input = re.ReplaceAllString(input, `$1 $2`)
	result := 1
	for _, v := range strings.Split(input, " ") {
		num, _ := strconv.Atoi(v)

		result *= num
	}

	return result
}

func mulAll(muls []string) int {
	sum := 0
	for _, m := range muls {
		re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
		m = re.ReplaceAllString(m, `$1 $2`)
		result := 1
		for _, v := range strings.Split(m, " ") {
			num, _ := strconv.Atoi(v)
			if num >= 1000 {
				continue
			}
			result *= num
		}
		sum += result

	}
	return sum
}

func main() {
	file, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	inputs := string(file)
	sum := 0
	sumEnabled := 0
	for _, input := range strings.Split(inputs, "\n") {
		r := renderer{line: input}
		muls := findCorrectInstructions(input)
		enabledMuls := findEnabledInstructions(input)
		r.render(enabledMuls)

		sum += mulAll(muls)
		sumEnabled += mulAll(enabledMuls)
	}
	log.Println(sum)
	log.Println(sumEnabled)
}
