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
	scratchCards := make(map[uint64]uint64)
	for scanner.Scan() {
		scratchCardWinnings(scanner.Text(), &scratchCards)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for _, cardCount := range scratchCards {
		result_sum += cardCount
	}
	fmt.Println(scratchCards)
	return result_sum
}

func scratchCardWinnings(line string, scratchCards *map[uint64]uint64) {
	card := strings.Split(line, ":")
	cardNum, err := strconv.ParseUint(strings.TrimLeft(card[0][5:], " "), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	(*scratchCards)[cardNum]++

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
	var cardsWon uint64 = 0
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
				cardsWon++
			}
		}
	}
	var i uint64
	for i = 1; i <= cardsWon; i++ {
		(*scratchCards)[cardNum+i] += (*scratchCards)[cardNum]
	}
}
