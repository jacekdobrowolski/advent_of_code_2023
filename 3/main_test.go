package main

import (
	"log"
	"os"
	"testing"
)

func TestEnginePartNumSum(t *testing.T) {
	file, err := os.ReadFile("test_input")
	if err != nil {
		log.Fatal(err)
	}

	result_sum := enginePartNumSum(file)
	if result_sum != 467835 {
		t.Fatalf("expected %d, got: %d", 467835, result_sum)
	}
}

func BenchmarkEnginePartNumSum(b *testing.B) {
	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		enginePartNumSum(file)
	}
}
