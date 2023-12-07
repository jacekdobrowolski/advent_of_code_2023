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

	result_sum := scratchCardWinningsSum(file)
	fmt.Printf("Result: %d\n", result_sum)
}

func scratchCardWinningsSum(file *os.File) uint64 {
	scanner := bufio.NewScanner(file)
	var result_sum uint64
	for scanner.Scan() {
		result_sum += scratchCardWinnings(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result_sum
}

func scratchCardWinnings(line string) uint64 {
	var points uint64 = 0
	card := strings.Split(line, ":")
	winningNumsAndElfsNums := strings.Split(card[1], "|")
	winningNumsStr, ElfsNumsStr := winningNumsAndElfsNums[0], winningNumsAndElfsNums[1]
	winningNums := make([]uint64, 0, 5)
	for _, winningNum := range strings.Split(winningNumsStr, " ") {
		if winningNum == "" {
			continue
		}
		num, err := strconv.ParseUint(winningNum, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		winningNums = append(winningNums, num)
	}
	for _, elfsNum := range strings.Split(ElfsNumsStr, " ") {
		if elfsNum == "" {
			continue
		}
		elfsNum, err := strconv.ParseUint(elfsNum, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		for _, winningNum := range winningNums {
			if elfsNum == winningNum {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
	}
	return points
}
