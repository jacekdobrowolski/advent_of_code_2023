package main

import (
	"os"
	"testing"
)

func TestPossibleGamesIDSum(t *testing.T) {
	file, err := os.Open("test_input")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	result_sum := possibleGamesIDSum(file)
	if result_sum != 2286 {
		t.Fatalf("expected %d, got: %d", 2286, result_sum)
	}
}

func BenchmarkPossibleGamesIDSum(b *testing.B) {
	file, err := os.Open("input")
	if err != nil {
		b.Fatal(err)
	}
	defer file.Close()

	for i := 0; i < b.N; i++ {
		possibleGamesIDSum(file)
	}
}
