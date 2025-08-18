package main

import (
	"fmt"
	"math"
)

// Method 1: Simple if-else for integers
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Method 2: Simple if-else for float64
func maxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// Method 3: Using math.Max (works only with float64)
func maxUsingMath(a, b float64) float64 {
	return math.Max(a, b)
}

// Method 4: Generic function (Go 1.18+) - works with any comparable numeric type
func Max[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Method 5: Using ternary-like approach with anonymous function
func maxTernary(a, b int) int {
	return func() int {
		if a > b {
			return a
		}
		return b
	}()
}

// Method 6: Max function for multiple integers using variadic parameters
func maxMultiple(nums ...int) int {
	if len(nums) == 0 {
		panic("at least one number required")
	}
	
	max := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}
	return max
}

// Method 7: Max with custom comparison function
func maxWithComparator[T any](a, b T, compare func(T, T) bool) T {
	if compare(a, b) {
		return a
	}
	return b
}

func demonstrateMaxFunctions() {
	fmt.Println("=== Max Function Demonstrations ===\n")

	// Method 1: Simple integers
	fmt.Printf("maxInt(10, 20): %d\n", maxInt(10, 20))
	fmt.Printf("maxInt(25, 15): %d\n", maxInt(25, 15))

	// Method 2: Floats
	fmt.Printf("maxFloat(10.5, 20.3): %.2f\n", maxFloat(10.5, 20.3))
	fmt.Printf("maxFloat(25.7, 15.2): %.2f\n", maxFloat(25.7, 15.2))

	// Method 3: Using math.Max
	fmt.Printf("maxUsingMath(10.5, 20.3): %.2f\n", maxUsingMath(10.5, 20.3))

	// Method 4: Generic function
	fmt.Printf("Max[int](10, 20): %d\n", Max[int](10, 20))
	fmt.Printf("Max[float64](10.5, 20.3): %.2f\n", Max[float64](10.5, 20.3))
	fmt.Printf("Max[int32](100, 200): %d\n", Max[int32](100, 200))

	// Method 5: Ternary-like
	fmt.Printf("maxTernary(10, 20): %d\n", maxTernary(10, 20))

	// Method 6: Multiple numbers
	fmt.Printf("maxMultiple(5, 15, 25, 10, 30): %d\n", maxMultiple(5, 15, 25, 10, 30))
	fmt.Printf("maxMultiple(100): %d\n", maxMultiple(100))

	// Method 7: With custom comparator
	isGreater := func(a, b int) bool { return a > b }
	fmt.Printf("maxWithComparator(10, 20, isGreater): %d\n", maxWithComparator(10, 20, isGreater))

	fmt.Println("\n=== Performance Comparison ===")
	// Simple benchmark-style comparison
	a, b := 1000000, 999999
	
	fmt.Printf("Testing with a=%d, b=%d\n", a, b)
	fmt.Printf("maxInt result: %d\n", maxInt(a, b))
	fmt.Printf("Max[int] result: %d\n", Max[int](a, b))
	fmt.Printf("math.Max result: %.0f\n", math.Max(float64(a), float64(b)))
}

// Example integration with your worker pool pattern
type JobResult struct {
	WorkerID int
	JobID    int
	Value    int
}

func findMaxResult(results []JobResult) JobResult {
	if len(results) == 0 {
		return JobResult{}
	}
	
	maxResult := results[0]
	for _, result := range results[1:] {
		if result.Value > maxResult.Value {
			maxResult = result
		}
	}
	return maxResult
}

func workerPoolMaxExample() {
	fmt.Println("\n=== Worker Pool Max Example ===")
	
	results := []JobResult{
		{WorkerID: 1, JobID: 1, Value: 42},
		{WorkerID: 2, JobID: 2, Value: 87},
		{WorkerID: 3, JobID: 3, Value: 23},
		{WorkerID: 1, JobID: 4, Value: 156},
		{WorkerID: 2, JobID: 5, Value: 91},
	}
	
	maxResult := findMaxResult(results)
	fmt.Printf("Max result: Worker %d, Job %d, Value %d\n", 
		maxResult.WorkerID, maxResult.JobID, maxResult.Value)
}

func runMaxDemo() {
	demonstrateMaxFunctions()
	workerPoolMaxExample()
}

// Uncomment and rename to main() to run this demo
// func main() {
//     runMaxDemo()
// }
