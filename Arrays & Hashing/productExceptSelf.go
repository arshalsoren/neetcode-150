// Given an integer array nums, return an array output where output[i] is the product of all the elements of nums except nums[i].
// Each product is guaranteed to fit in a 32-bit integer.
// https://leetcode.com/problems/products-of-array-except-self/

package main

import (
	"fmt"
)

func productExceptSelfV2(nums []int) []int {
	res := make([]int, len(nums))
	prefix := 1
	for i := 0; i < len(nums); i++ {
		res[i] = prefix
		prefix *= nums[i]
	}
	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}
	return res
}

func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		product := 1
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			product *= nums[j]
		}
		output[i] = product
	}
	return output
}

func main() {
	nums := []int{1, 2, 4, 6}
	fmt.Println(productExceptSelf(nums))
	fmt.Println(productExceptSelfV2(nums))
}
