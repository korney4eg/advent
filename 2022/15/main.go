package main

import (
	"advent/utils"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func parseCommand(line string) *Sensor {
	sensor := &Sensor{}
	var re = regexp.MustCompile(`\d+`)
	results := re.FindAllString(line, -1)
	sensor.X, _ = strconv.ParseInt(results[0], 10, 64)
	sensor.Y, _ = strconv.ParseInt(results[1], 10, 64)
	sensor.BeaconX, _ = strconv.ParseInt(results[2], 10, 64)
	sensor.BeaconY, _ = strconv.ParseInt(results[3], 10, 64)
	sensor.Distance = int64(math.Abs(float64(sensor.X-sensor.BeaconX)) + math.Abs(float64(sensor.Y-sensor.BeaconY)))
	return sensor
}
func isSensorWithinDistance(x, y int64, sensor *Sensor) bool {
	return getManhattanDistance(x, y, sensor.X, sensor.Y) <= sensor.Distance
}

func getManhattanDistance(x1, y1, x2, y2 int64) int64 {
	dx := int64(math.Abs(float64(x1 - x2)))
	dy := int64(math.Abs(float64(y1 - y2)))
	return dx + dy
}

type Sensor struct {
	X        int64
	Y        int64
	Distance int64
	BeaconX  int64
	BeaconY  int64
}

// func getPoint(x,y, sensors []*Sensor)string{
// 	for _, sensor:= range sensors{
// 	}
// 	return ""
// }

func draw(minX, minY, maxX, maxY int64, sensors []*Sensor) {
	output := [][]string{}
	for y := minY; y <= maxY; y++ {
		line := []string{}
		for x := minX; x <= maxX; x++ {
			line = append(line, ".")
		}
		output = append(output, line)
	}

	// single sensor:
	for _, sensor := range sensors {
		for ny := sensor.Y - sensor.Distance; ny <= sensor.Y+sensor.Distance; ny++ {
			y := ny - minY
			// y := -minY + ny
			for nx := sensor.X - sensor.Distance; nx <= sensor.X+sensor.Distance; nx++ {
				// x := -minX + nx
				x := nx - minX
				// if x < minX {
				// 	continue
				// }
				// if x > maxX {
				// 	break
				// }
				if nx == sensor.X && ny == sensor.Y {
					output[y][x] = "S"
				} else if nx == sensor.BeaconX && ny == sensor.BeaconY {
					output[y][x] = "b"
				} else if output[y][x] == "." && getManhattanDistance(nx, ny, sensor.X, sensor.Y) <= sensor.Distance {
					output[y][x] = "#"
				}
			}
		}
	}
	outputStr := "\n"
	for i, line := range output {
		outputStr += fmt.Sprintf("%s|%d\n", strings.Join(line, ""), int64(i)+minY)
	}
	log.Println(outputStr)

}

func main() {
	lines := utils.ReadFileToStringsList("2022/15/test.txt")
	sensors := []*Sensor{}
	var minX, maxX, minY, maxY, curY int64
	minX = 999999
	maxX = -9999999
	minY = 999999
	maxY = -9999999
	curY = 10
	for _, line := range lines {
		newSensor := parseCommand(line)
		minY = int64(math.Min(float64(minY), float64(newSensor.Y-newSensor.Distance)))
		maxY = int64(math.Max(float64(maxY), float64(newSensor.Y+newSensor.Distance)))
		minX = int64(math.Min(float64(minX), float64(newSensor.X-newSensor.Distance)))
		maxX = int64(math.Max(float64(maxX), float64(newSensor.X+newSensor.Distance)))
		sensors = append(sensors, newSensor)
	}
	foundPoints := 0
	draw(minX, minY, maxX, maxY, sensors)
	for x := minX; x <= maxX; x++ {
		for _, sensor := range sensors {
			if isSensorWithinDistance(x, curY, sensor) {
				foundPoints++
				break
			}
		}

	}
	log.Println(foundPoints - 1)

}
