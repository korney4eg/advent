package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type report struct {
	levels []int
}

func (r *report) isLevelSafe() (safe bool, ops int) {
	sign := 0
	newSign := 0
	val := 0
	levels := r.levels
	for i := 0; i < len(levels)-1; i++ {
		if levels[i] < levels[i+1] {
			newSign = 1
		} else if levels[i] > levels[i+1] {
			newSign = -1
		} else {
			// log.Println(printReport(levels, []int{levels[i], levels[i+1]}), "level wasnt't changed")
			return safe, ops
		}
		ops++
		val = int(math.Abs(float64(levels[i] - levels[i+1])))
		if (sign != 0 && sign == newSign) || sign == 0 {
			sign = newSign
		} else {
			// log.Println("sing changed: sign", sign, "newSign", newSign)
			return safe, ops
		}
		ops++
		if val > 3 {
			ops++
			// log.Println("level difference is less than 3 =", val)
			return safe, ops
		}
	}
	safe = true
	return safe, ops
}

func (r *report) printReport() string {
	retVal := ""
	for _, level := range r.levels {
		val := strconv.Itoa(level)
		// if slices.Contains(dampers, i) {
		// 	val = "(" + val + ")"
		// }

		retVal += val + " "
	}
	retVal = fmt.Sprintf("[ %s]", retVal)
	return retVal
}

func countSafeLevels(reports []report) (int, int) {
	sum := 0
	ops := 0
	for _, rep := range reports {
		ops++
		ok, newOps := rep.isLevelSafe()
		ops += newOps
		if ok {
			log.Printf("✅ %s", rep.printReport())
			sum += 1
			continue
		} else {
			log.Printf("❌ %s", rep.printReport())
		}

	}
	return sum, ops
}

func (r *report) isLevelSafeWithDamper() (safe bool, ops int) {
	ok, newOps := r.isLevelSafe()
	if ok {
		return ok, newOps
	}
	levels := r.levels
	for i := 0; i < len(levels); i++ {
		levelSet := make([]int, 0)
		for j, level := range levels {
			if j != i {
				levelSet = append(levelSet, level)
			}
		}
		rr := report{levels: levelSet}
		log.Println("levelSet", levelSet)
		ok, newOps := rr.isLevelSafe()
		ops += newOps
		if ok {
			return ok, newOps
		}
		ops++
	}
	safe = false
	return safe, ops
}
func countSafeLevelsWithDamper(reports []report) (sum int, ops int) {
	for _, rep := range reports {
		ops++
		ok, newOps := rep.isLevelSafeWithDamper()
		ops += newOps
		if ok {
			sum += 1
			continue
		} else {
			log.Printf("is not safe")
		}

	}
	return sum, ops
}

func getInput(input string) (reports []report) {
	for _, line := range strings.Split(input, "\n") {
		rep := report{levels: []int{}}
		if line == "" {
			continue
		}
		for _, field := range strings.Split(line, " ") {
			level, _ := strconv.Atoi(field)
			rep.levels = append(rep.levels, level)
		}
		reports = append(reports, rep)

	}
	return reports

}

func main() {
	file, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	input := string(file)
	reports := getInput(input)
	// countSafe, ops := countSafeLevels(reports)
	// log.Printf("countSafe %d, in %d operations", countSafe, ops)
	countSafeWithDampers, ops := countSafeLevelsWithDamper(reports)
	log.Printf("countSafe with dampers %d, in %d operations", countSafeWithDampers, ops)
}
