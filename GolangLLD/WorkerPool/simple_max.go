package main

import (
	"fmt"
	"math"
)

// Simple max function for integers
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Generic max function (Go 1.18+)
func Max[T ~int | ~float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Method 1: Simple if-else
	fmt.Println("=== Simple Max Functions ===")
	fmt.Printf("maxInt(15, 25): %d\n", maxInt(15, 25))
	fmt.Printf("maxInt(100, 50): %d\n", maxInt(100, 50))
	
	// Method 2: Using math.Max (for floats)
	fmt.Println("\n=== Using math.Max ===")
	fmt.Printf("math.Max(15.5, 25.3): %.2f\n", math.Max(15.5, 25.3))
	fmt.Printf("math.Max(100.7, 50.2): %.2f\n", math.Max(100.7, 50.2))
	
	// Method 3: Generic function
	fmt.Println("\n=== Generic Max Function ===")
	fmt.Printf("Max[int](15, 25): %d\n", Max[int](15, 25))
	fmt.Printf("Max[float64](15.5, 25.3): %.2f\n", Max[float64](15.5, 25.3))
	
	// Practical examples
	fmt.Println("\n=== Practical Examples ===")
	
	// Example 1: Finding max in worker results
	workerResults := []int{45, 23, 67, 12, 89, 34}
	maxResult := workerResults[0]
	for _, result := range workerResults[1:] {
		maxResult = maxInt(maxResult, result)
	}
	fmt.Printf("Max worker result: %d\n", maxResult)
	
	// Example 2: Max between two job processing times
	jobTime1 := 125 // milliseconds
	jobTime2 := 87  // milliseconds
	maxTime := maxInt(jobTime1, jobTime2)
	fmt.Printf("Max processing time: %d ms\n", maxTime)
}
