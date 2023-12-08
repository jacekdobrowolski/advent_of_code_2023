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
	raceTimes := parseColumns(timeAndDistance[0])
	winningDistances := parseColumns(timeAndDistance[1])
	result := 1
	for raceIdx := 0; raceIdx < len(raceTimes); raceIdx++ {

		fmt.Printf("Race: %d, time: %d, recordDistance: %d\n", raceIdx, raceTimes[raceIdx], winningDistances[raceIdx])
		winningMargin := 0
		for buttonPressTime := 1; buttonPressTime < raceTimes[raceIdx]; buttonPressTime++ {
			speed := buttonPressTime
			travelTime := raceTimes[raceIdx] - buttonPressTime
			distance := speed * travelTime
			if distance > winningDistances[raceIdx] {
				winningMargin++
				fmt.Printf("Winning posible for press time %d with distance %d\n", buttonPressTime, distance)
			}
		}
		result *= winningMargin
	}

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
