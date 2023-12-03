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

	result_sum := enginePartNumSum(file)
	fmt.Printf("\nResult: %d\n", result_sum)
}

func enginePartNumSum(schematic_file []byte) uint64 {
	schematic := string(schematic_file)
	var result_sum uint64
	lines := strings.Split(schematic, "\n")
	for lineNum, line := range lines {
		index := 0
		for index < len(line) {
			if lineNum == 139 {
				fmt.Println("close")
			}
			partNumStartIdx := strings.IndexAny(line[index:], "0123456789")
			if partNumStartIdx < 0 {
				break
			}
			partNumStartIdx += index
			partNumEndIdx := partNumStartIdx
			for ; partNumEndIdx < len(line); partNumEndIdx++ {
				// fmt.Printf("line: %d, idx: %d, start: %d, end: %d\n", lineNum, index, partNumStartIdx, partNumEndIdx)
				if line[partNumEndIdx] < '0' || line[partNumEndIdx] > '9' {
					break
				}
			}
			index = partNumEndIdx
			result_sum += checkPartNumValid(partNumStartIdx, partNumEndIdx, lineNum, &lines)
		}
	}
	return result_sum
}

func checkPartNumValid(partNumStartIdx, partNumEndIdx int, lineNum int, lines *[]string) uint64 {

	signAreaTopLineNum := lineNum - 1
	if signAreaTopLineNum < 0 {
		signAreaTopLineNum = 0
	}
	signAreaBottomLineNum := lineNum + 2
	if signAreaBottomLineNum >= len(*lines) {
		signAreaBottomLineNum = len(*lines) - 1
	}
	signAreaStartIdx := partNumStartIdx - 1
	if signAreaStartIdx < 0 {
		signAreaStartIdx = 0
	}
	signAreaEndIdx := partNumEndIdx + 1
	if signAreaEndIdx >= len((*lines)[lineNum]) {
		signAreaEndIdx = len((*lines)[lineNum])
	}
	partNum, err := strconv.ParseUint((*lines)[lineNum][partNumStartIdx:partNumEndIdx], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	var signArea string
	for _, line := range (*lines)[signAreaTopLineNum:signAreaBottomLineNum] {
		signArea += line[signAreaStartIdx:signAreaEndIdx]
	}
	if strings.ContainsAny(signArea, "!@#$%^&*()_+{}|:\"<>?,/;[]\\`~'=-") {
		return partNum
	} else {
		return 0
	}
}
