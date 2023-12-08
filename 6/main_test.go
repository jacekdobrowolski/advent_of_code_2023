package main

import (
	"log"
	"os"
	"testing"
)

func TestWinningMargin(t *testing.T) {
	file, err := os.ReadFile("test_input")
	if err != nil {
		log.Fatal(err)
	}

	result_sum := winningMargin(file)
	if result_sum != 288 {
		t.Fatalf("expected %d, got: %d", 288, result_sum)
	}
}

func BenchmarkWinningMargin(b *testing.B) {
	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		winningMargin(file)
	}
}
