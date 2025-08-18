package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// two pointers appraoch to find the largest area here now

func largestArea(arr []int) int {
	start := 0
	end := len(arr) - 1
	area := 0

	for start < end {
		curr_area := (end - start) * min(arr[start], arr[end])
		area = max(area, curr_area)
		if arr[start] < arr[end] {
			start++
		} else {
			end--
		}
	}
	return area
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	ans := largestArea(arr)
	fmt.Println("Largest Area:", ans)
	return
}
