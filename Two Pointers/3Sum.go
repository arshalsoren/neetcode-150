// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] where nums[i] + nums[j] + nums[k] == 0, and the indices i, j and k are all distinct.
// The output should not contain any duplicate triplets. You may return the output and the triplets in any order.
// https://leetcode.com/problems/3sum/

package main

import (
	"fmt"
	"sort"
)

type Solution struct{}

func (s *Solution) threeSum(nums []int) [][]int {
	res := [][]int{}
	if len(nums) < 3 {
		return res
	}

	sort.Ints(nums)
	for index, a := range nums {
		if a > 0 {
			break
		}

		if index > 0 && a == nums[index-1] {
			continue
		}

		left := index + 1
		right := len(nums) - 1
		for left < right {
			threeSum := a + nums[left] + nums[right]
			if threeSum > 0 {
				right--
			} else if threeSum < 0 {
				left++
			} else {
				res = append(res, []int{a, nums[left], nums[right]})
				left++
				right--
				for nums[left] == nums[left-1] && left < right {
					left++
				}
			}
		}
	}

	return res
}

func main() {
	solution := &Solution{}
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("Test case 1: %v\n", solution.threeSum(nums))
}
