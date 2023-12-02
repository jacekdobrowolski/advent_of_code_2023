package main

import (
	"os"
	"testing"
)

func TestCallibrationSum(t *testing.T) {
	file, err := os.Open("input")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	result_sum := CallibrationSum(file)
	if result_sum != 54578 {
		t.Fatalf("expected %d, got: %d", 54578, result_sum)
	}
}

func BenchmarkCallibrationSum(b *testing.B) {
	file, err := os.Open("input")
	if err != nil {
		b.Fatal(err)
	}
	defer file.Close()

	for i := 0; i < b.N; i++ {
		CallibrationSum(file)
	}
}
