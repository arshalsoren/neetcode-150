// Given an integer array nums, return true if any value appears more than once in the array, otherwise return false.

package main

import "fmt"

type Solution struct{}

func (s *Solution) hasDuplicate(nums []int) bool {
	res := make(map[int]bool)
	for _, val := range nums {
		if _, exists := res[val]; exists {
			return true
		}
		res[val] = true
	}
	return false
}

func main() {
	solution := &Solution{}
	nums1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Test case 1: %v\n", solution.hasDuplicate(nums1))

	nums2 := []int{1, 2, 3, 3, 4}
	fmt.Printf("Test case 2: %v\n", solution.hasDuplicate(nums2))
}
