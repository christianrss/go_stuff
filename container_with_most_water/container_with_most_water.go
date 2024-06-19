package main

import "fmt"

// Time complexity: O(N^2)
// Space complexity O(1)
// Brute Force Solution
func maxArea(heights []int) int {
	var max_area int = 0
	var n int = len(heights)
	for p1 := 0; p1 < n; p1++ {

		for p2 := p1 + 1; p2 < n; p2++ {
			var length int = min(heights[p1], heights[p2])
			var width int = p2 - p1
			var area int = length * width
			max_area = max(max_area, area)
		}
	}

	return max_area
}

func main() {
	var max_area int = maxArea([]int{5, 9, 2, 4})
	fmt.Printf("Max area: %d \n", max_area)
}
