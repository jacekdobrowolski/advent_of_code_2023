package main

import (
	"log"
	"os"
	"testing"
)

func TestLowestLocation(t *testing.T) {
	file, err := os.ReadFile("test_input")
	if err != nil {
		log.Fatal(err)
	}

	result_sum := lowestLocation(file)
	if result_sum != 46 {
		t.Fatalf("expected %d, got: %d", 46, result_sum)
	}
}

func BenchmarkLowestLocation(b *testing.B) {
	file, err := os.ReadFile("test_input")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		lowestLocation(file)
	}
}
