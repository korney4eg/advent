package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type test struct {
	isDivisibleBy uint64
	trueMonkeyID  int
	falseMonkeyID int
}

type Monkey struct {
	inspections int
	id          int
	items       []uint64
	operation   string
	test        test
}

func getMonkeyID(line string) int {
	monkeyIdString := strings.TrimPrefix(line, "Monkey ")
	monkeyIdString = strings.TrimSuffix(monkeyIdString, ":")
	monkeyId, _ := strconv.Atoi(monkeyIdString)
	return monkeyId

}

func getStartingItems(line string) (items []uint64) {
	itemsLine := strings.TrimPrefix(line, "  Starting items: ")
	itemsList := strings.Split(itemsLine, ", ")
	for _, item := range itemsList {
		itemInt, _ := strconv.ParseInt(item, 10, 64)
		items = append(items, uint64(itemInt))
	}
	return items
}

func getOperation(line string) string {
	return strings.TrimPrefix(line, "  Operation: new = old ")
}

func getTest(lines []string) test {
	t := test{}
	testString := strings.TrimPrefix(lines[0], "  Test: divisible by ")
	isDivisibleBy, err := strconv.ParseUint(testString, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	t.isDivisibleBy = isDivisibleBy
	trueMonkeyIDString := strings.TrimPrefix(lines[1], "    If true: throw to monkey ")
	t.trueMonkeyID, err = strconv.Atoi(trueMonkeyIDString)
	if err != nil {
		log.Fatal(err)
	}
	falseMonkeyIDString := strings.TrimPrefix(lines[2], "    If false: throw to monkey ")
	t.falseMonkeyID, err = strconv.Atoi(falseMonkeyIDString)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func (m *Monkey) readFromString(data string) {
	lines := strings.Split(data, "\n")
	m.id = getMonkeyID(lines[0])
	m.items = getStartingItems(lines[1])
	m.operation = getOperation(lines[2])
	m.test = getTest([]string{lines[3], lines[4], lines[5]})
}

func getMonkeybyID(monkeys []*Monkey, id int) *Monkey {
	for _, monkey := range monkeys {
		if monkey.id == id {
			return monkey
		}
	}
	return nil
}

func (m *Monkey) getWorryLevelAfteroperation(worryLevel uint64) uint64 {
	var secondNumber uint64
	operations := strings.Split(m.operation, " ")
	if operations[1] == "old" {
		secondNumber = worryLevel
	} else {
		secondNumber, _ = strconv.ParseUint(operations[1], 10, 64)
	}
	switch operations[0] {
	case "+":
		return worryLevel + secondNumber
	case "-":
		return worryLevel - secondNumber
	case "*":
		return worryLevel * secondNumber
	case "/":
		return worryLevel / secondNumber
	}
	return 0
}

func (m *Monkey) choseMonkeyIDToThrowItem(worryLevel uint64) int {
	if worryLevel%m.test.isDivisibleBy == 0 {
		// log.Printf("    Current worry level is divisible by %d.", m.test.isDivisibleBy)
		// log.Printf("    Item with worry level %d is thrown to monkey %d.", worryLevel, m.test.trueMonkeyID)
		return m.test.trueMonkeyID
	}
	// log.Printf("    Current worry level is not divisible by %d.", m.test.isDivisibleBy)
	// log.Printf("    Item with worry level %d is thrown to monkey %d.", worryLevel, m.test.falseMonkeyID)
	return m.test.falseMonkeyID
}

func (m *Monkey) throwItemToMonkey(itemWorryLevel uint64, monkey *Monkey) {
	monkey.items = append(monkey.items, itemWorryLevel)
}

func playRound(monkeys []*Monkey) {
	for _, monkey := range monkeys {
		// log.Printf("Monkey %d:\n", monkey.id)
		for _, item := range monkey.items {
			// log.Printf("  Monkey inspects an item with a worry level of %d.\n", item)
			worryLevelAfterOperation := monkey.getWorryLevelAfteroperation(item)
			// log.Printf("    Worry level is '%s' to %d.", monkey.operation, worryLevelAfterOperation)
			// worryLevelAfterBoared := worryLevelAfterOperation / 3
			worryLevelAfterBoared := worryLevelAfterOperation
			// log.Printf("    Monkey gets bored with item. Worry level is divided by 3 to %d.", worryLevelAfterBoared)
			monkeyIDToThrow := monkey.choseMonkeyIDToThrowItem(worryLevelAfterBoared)
			monkeyToThrow := getMonkeybyID(monkeys, monkeyIDToThrow)
			monkey.throwItemToMonkey(worryLevelAfterBoared, monkeyToThrow)

		}
		monkey.inspections += len(monkey.items)
		monkey.items = []uint64{}
	}

}
func monkeysOutput(monkeys []*Monkey) string {
	output := ""
	for _, monkey := range monkeys {
		output += fmt.Sprintf("Monkey %d inspected items %d times:%v\n", monkey.id, monkey.inspections, monkey.items)
	}
	return output

}

func getMonkeyBusiness(monkeys []*Monkey) int {
	inspections := []int{}
	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.inspections)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))

	return inspections[0] * inspections[1]

}

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	file, err := os.ReadFile("test.txt")

	if err != nil {
		log.Fatal(err)
	}
	monkeysInput := strings.Split(string(file), "\n\n")
	monkeys := []*Monkey{}
	for _, monkeyString := range monkeysInput {
		monkey := &Monkey{}
		monkey.readFromString(monkeyString)
		monkeys = append(monkeys, monkey)
	}
	for i := 0; i < 10000; i++ {
		playRound(monkeys)
		if i < 25 || (i+1)%1000 == 0 {
			log.Printf("== After round %d ==", i+1)
			log.Println(monkeysOutput(monkeys))
		}
	}
	log.Println(getMonkeyBusiness(monkeys))

}
