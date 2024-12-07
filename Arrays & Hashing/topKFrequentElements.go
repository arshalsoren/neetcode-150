// Given an integer array nums and an integer k, return the k most frequent elements within the array.
// The test cases are generated such that the answer is always unique.
// You may return the output in any order.
// https://leetcode.com/problems/top-k-frequent-elements/

package main

import (
	"fmt"
)

type Solution struct {
}

func (s *Solution) TopKFrequent(nums []int, k int) []int {
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num]++
	}

	result := make([]int, 0)
	for key, value := range frequency {
		if value >= k {
			result = append(result, key)
		}
	}

	return result
}

func main() {
	nums := []int{1, 2, 2, 2, 3, 3}
	k := 2
	solution := Solution{}
	result := solution.TopKFrequent(nums, k)
	fmt.Println(result)
}
