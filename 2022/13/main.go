package main

import (
	"advent/2022/13/signals"
	"advent/utils"
	"log"
	"sort"
)

type packet struct {
	text string
}

type byText []packet

func (a byText) Len() int           { return len(a) }
func (a byText) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byText) Less(i, j int) bool { return signals.IsRightOrder(a[i].text, a[j].text) }

func main() {
	lines := utils.ReadFileToStringsList("2022/13/input.txt")
	packets := []packet{}
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		newPacket := packet{text: lines[i]}
		packets = append(packets, newPacket)

	}
	dividerPackets := []packet{{text: "[[2]]"}, {text: "[[6]]"}}
	packets = append(packets, dividerPackets...)
	sort.Sort(byText(packets))
	result := 1
	for i, pack := range packets {
		log.Println(pack)
		if pack.text == "[[2]]" || pack.text == "[[6]]" {
			result *= (i + 1)
		}
	}
	log.Println(result)

}
