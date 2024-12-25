// You are given an integer array heights where heights[i] represents the height of the ith bar.
// You may choose any two bars to form a container. Return the maximum amount of water a container can store.
// https://leetcode.com/problems/container-with-most-water/

package main

import (
	"fmt"
)

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}
	left := 0
	right := len(height) - 1
	max_area := 0
	for left < right {
		max_area = max(max_area, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}
	}
	return max_area
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	// Output: 49
	fmt.Println(maxArea(height))
}
