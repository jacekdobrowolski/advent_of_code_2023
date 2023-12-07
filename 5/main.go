package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	result_sum := lowestLocation(file)
	fmt.Printf("\nResult: %d\n", result_sum)
}

type Mapping struct {
	source      uint64
	destination uint64
	offset      uint64
}

func lowestLocation(fileBytes []byte) uint64 {
	fileStr := string(fileBytes)
	seedsAndMaps := strings.Split(fileStr, "\n\n")
	seeds := parseSeeds(seedsAndMaps[0])
	maps := make([][]Mapping, 7)
	for i, offsetMappingsStr := range seedsAndMaps[1:] {
		maps[i] = *parseMappings(offsetMappingsStr)
	}

	minLocation := ^uint64(0)
	for _, seed := range seeds {
		x := seed
		for _, mappings := range maps {
			x = resolveMapping(x, mappings)
		}
		if x < minLocation {
			minLocation = x
		}
	}
	return minLocation
}

func parseSeeds(s string) []uint64 {
	seedsStr := strings.SplitAfter(s, " ")[1:]
	seeds := make([]uint64, len(seedsStr))
	for i, seedStr := range seedsStr {
		seed, err := strconv.ParseUint(strings.Trim(seedStr, " "), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		seeds[i] = seed
	}
	return seeds
}

func parseMappings(offsetMap string) *[]Mapping {
	mappings := make([]Mapping, 0)
	for _, line := range strings.Split(offsetMap, "\n")[1:] {
		mappingsStrSlice := strings.Split(line, " ")
		for i, elem := range mappingsStrSlice {
			mappingsStrSlice[i] = strings.Trim(elem, " ")
		}
		destination, err := strconv.ParseUint(mappingsStrSlice[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		source, err := strconv.ParseUint(mappingsStrSlice[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		offset, err := strconv.ParseUint(mappingsStrSlice[2], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		mappings = append(mappings, Mapping{source: source, destination: destination, offset: offset})
	}
	slices.SortFunc(mappings,
		func(a, b Mapping) int {
			return cmp.Compare(a.source, b.source)
		})
	return &mappings
}

func resolveMapping(x uint64, mappings []Mapping) uint64 {
	for _, mapping := range mappings {
		if x >= mapping.source && x < mapping.source+mapping.offset {
			return (x + mapping.destination) - mapping.source
		}
	}
	return x
}
