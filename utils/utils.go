package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFileToStringsList(filName string) []string {
	file, err := os.Open(filName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	stringsList := []string{}
	for scanner.Scan() {
		stringsList = append(stringsList, scanner.Text())
	}
	return stringsList
}
