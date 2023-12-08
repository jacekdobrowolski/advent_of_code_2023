package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	result_sum := winningMargin(file)
	fmt.Printf("\nResult: %d\n", result_sum)
}

func winningMargin(fileBytes []byte) int {
	fileStr := string(fileBytes)
	timeAndDistance := strings.Split(fileStr, "\n")
	raceTime := parseBadKerning(timeAndDistance[0])
	winningDistance := parseBadKerning(timeAndDistance[1])
	result := 1

	fmt.Printf("time: %d, recordDistance: %d\n", raceTime, winningDistance)
	var minPressTime int
	for buttonPressTime := 1; buttonPressTime < raceTime; buttonPressTime++ {
		speed := buttonPressTime
		travelTime := raceTime - buttonPressTime
		distance := speed * travelTime
		if distance > winningDistance {
			minPressTime = buttonPressTime
			fmt.Printf("Minimum press time to win %d\n", buttonPressTime)
			break
		}
	}
	var maxPressTime int
	for buttonPressTime := raceTime; buttonPressTime > 1; buttonPressTime-- {
		speed := buttonPressTime
		travelTime := raceTime - buttonPressTime
		distance := speed * travelTime
		if distance > winningDistance {
			maxPressTime = buttonPressTime
			fmt.Printf("Maximum press time to win %d\n", buttonPressTime)
			break
		}
	}
	result *= maxPressTime - minPressTime + 1

	return result
}

func parseColumns(columnsLine string) []int {
	columnsStr := strings.SplitAfter(columnsLine, ":")[1]
	columnsStrSlice := strings.Split(columnsStr, " ")
	columns := make([]int, 0)
	for _, raceTimeStr := range columnsStrSlice {
		columnStrTrimmed := strings.Trim(raceTimeStr, " ")
		if columnStrTrimmed == "" {
			continue
		}
		column, err := strconv.Atoi(columnStrTrimmed)
		if err != nil {
			log.Fatal(err)
		}
		columns = append(columns, column)
	}
	return columns
}

func parseBadKerning(line string) int {
	numbers := strings.SplitAfter(line, ":")[1]
	numberWithoutSpaces := strings.Replace(numbers, " ", "", -1)
	number, err := strconv.Atoi(numberWithoutSpaces)
	if err != nil {
		log.Fatal(err)
	}
	return number
}
