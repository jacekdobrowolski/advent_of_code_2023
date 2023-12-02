package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result_sum := possibleGamesIDSum(file)
	fmt.Printf("Result: %d\n", result_sum)
}

func possibleGamesIDSum(file *os.File) uint64 {
	bag := map[string]uint64{
		"red":   12,
		"green": 13,
		"blue":  14}

	scanner := bufio.NewScanner(file)
	var result_sum uint64
	for scanner.Scan() {
		result_sum += possibleGameID(scanner.Text(), bag)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result_sum
}

func possibleGameID(line string, bag map[string]uint64) uint64 {
	line_split_colon := strings.Split(line, ":")
	gameID, err := strconv.ParseUint(line_split_colon[0][5:], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	rounds := strings.Split(line_split_colon[1], ";")
	for _, round := range rounds {
		handful := strings.Split(round, ",")
		for _, cubes := range handful {
			countAndColor := strings.Split(cubes, " ")
			cubeCountStr, cubeColor := countAndColor[1], countAndColor[2]
			cubeCount, err := strconv.ParseUint(cubeCountStr, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			if cubesInBag, ok := bag[cubeColor]; ok {
				if cubeCount > cubesInBag {
					return 0
				}
			}
		}
	}
	return gameID
}
