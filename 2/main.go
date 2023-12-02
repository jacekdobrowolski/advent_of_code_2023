package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type gameBag map[string]uint64

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
	bag := gameBag{
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

func possibleGameID(line string, bag gameBag) uint64 {
	line_split_colon := strings.Split(line, ":")
	// gameID, err := strconv.ParseUint(line_split_colon[0][5:], 10, 64)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	rounds := strings.Split(line_split_colon[1], ";")
	minimumBag := gameBag{
		"red":   0,
		"green": 0,
		"blue":  0}
	for _, round := range rounds {
		handful := strings.Split(round, ",")
		for _, cubes := range handful {
			countAndColor := strings.Split(cubes, " ")
			cubeCountStr, cubeColor := countAndColor[1], countAndColor[2]
			cubeCount, err := strconv.ParseUint(cubeCountStr, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			if cubesInBag, ok := minimumBag[cubeColor]; ok {
				if cubeCount > cubesInBag {
					minimumBag[cubeColor] = cubeCount
				}
			}
		}
	}
	return powerOfGame(minimumBag)
}

func powerOfGame(bag gameBag) uint64 {
	var power uint64 = 1
	for _, count := range bag {
		power *= count
	}
	return power
}
