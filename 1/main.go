package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var result_sum int
	for scanner.Scan() {
		fmt.Print(scanner.Text())
		number := findNumber(scanner.Text())
		result_sum += number
		fmt.Println(number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: %d\n", result_sum)
}

type callibrationNumber struct {
	number int
	index  int
}

func newCallibrationNumber(number int, index int) callibrationNumber {
	return callibrationNumber{number: number, index: index}
}

func addSpelledCallibrationNumber(word string, number int, spelling string, numbers *[]callibrationNumber) {
	index := strings.Index(word, spelling)
	if index > -1 {
		*numbers = append(*numbers, newCallibrationNumber(number, index))
	}
	index = strings.LastIndex(word, spelling)
	if index > -1 {
		*numbers = append(*numbers, newCallibrationNumber(number, index))
	}
}

func findNumber(word string) int {
	numbers := make([]callibrationNumber, 0, 2)
	addSpelledCallibrationNumber(word, 1, "one", &numbers)
	addSpelledCallibrationNumber(word, 2, "two", &numbers)
	addSpelledCallibrationNumber(word, 3, "three", &numbers)
	addSpelledCallibrationNumber(word, 4, "four", &numbers)
	addSpelledCallibrationNumber(word, 5, "five", &numbers)
	addSpelledCallibrationNumber(word, 6, "six", &numbers)
	addSpelledCallibrationNumber(word, 7, "seven", &numbers)
	addSpelledCallibrationNumber(word, 8, "eight", &numbers)
	addSpelledCallibrationNumber(word, 9, "nine", &numbers)

	for i := 0; i < len(word); i++ {
		if word[i] >= '0' && word[i] <= '9' {
			numbers = append(numbers, newCallibrationNumber(int(word[i])-int('0'), i))
			break
		}
	}
	for i := len(word) - 1; i >= 0; i-- {
		if word[i] >= '0' && word[i] <= '9' {
			numbers = append(numbers, newCallibrationNumber(int(word[i])-int('0'), i))
			break
		}
	}
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i].index < numbers[j].index
	})
	fmt.Println(numbers)
	firstNumber := numbers[0].number
	lastnumber := numbers[len(numbers)-1].number
	return firstNumber*10 + lastnumber
}
