// Given an array of integers numbers that is sorted in non-decreasing order.
// Return the indices (1-indexed) of two numbers, [index1, index2], such that they add up to a given target number target and index1 < index2. Note that index1 and index2 cannot be equal, therefore you may not use the same element twice.
// You may assume that each input would have exactly one solution and you may not use the same element twice.

package main

import "fmt"

type Solution struct{}

func (s *Solution) twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	return []int{}
}

func main() {
	solution := &Solution{}
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("Test case 1: %v\n", solution.twoSum(nums, target))

	nums = []int{2, 3, 4}
	target = 6
	fmt.Printf("Test case 2: %v\n", solution.twoSum(nums, target))

	nums = []int{-1, 0}
	target = -1
	fmt.Printf("Test case 3: %v\n", solution.twoSum(nums, target))
}
