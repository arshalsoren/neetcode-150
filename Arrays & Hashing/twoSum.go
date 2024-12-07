// Given an array of integers nums and an integer target, return the indices i and j such that nums[i] + nums[j] == target and i != j.
// You may assume that every input has exactly one pair of indices i and j that satisfy the condition.
// Return the answer with the smaller index first.
// https://leetcode.com/problems/two-sum/

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	diffPosMap := make(map[int]int)
	for pos, val := range nums {
		diff := target - val
		if _, ok := diffPosMap[val]; ok {
			return []int{diffPosMap[val], pos}
		}
		diffPosMap[diff] = pos
	}
	return []int{}
}

func main() {
	nums := []int{3, 4, 5, 6}
	target := 7
	fmt.Println(twoSum(nums, target))
	nums = []int{4, 5, 6}
	target = 10
	fmt.Println(twoSum(nums, target))
	nums = []int{5, 5}
	target = 10
	fmt.Println(twoSum(nums, target))
}
