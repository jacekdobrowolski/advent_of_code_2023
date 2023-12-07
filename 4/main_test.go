package main

import (
	"os"
	"testing"
)

func TestScratchCardWinningsSum(t *testing.T) {
	file, err := os.Open("test_input")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	result_sum := scratchCardWinningsSum(file)
	if result_sum != 30 {
		t.Fatalf("expected %d, got: %d", 30, result_sum)
	}
}

func BenchmarkScratchCardWinningsSum(b *testing.B) {
	file, err := os.Open("input")
	if err != nil {
		b.Fatal(err)
	}
	defer file.Close()

	for i := 0; i < b.N; i++ {
		scratchCardWinningsSum(file)
	}
}
