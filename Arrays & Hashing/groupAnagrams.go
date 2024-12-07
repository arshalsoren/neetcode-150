// Given an array of strings strs, group all anagrams together into sublists. You may return the output in any order.
// An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.
// leetcode.com/problems/group-anagrams

package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, str := range strs {
		key := sortString(str)
		if _, ok := anagramMap[key]; !ok {
			anagramMap[key] = []string{str}
		} else {
			anagramMap[key] = append(anagramMap[key], str)
		}
	}

	var result [][]string
	for _, v := range anagramMap {
		result = append(result, v)
	}
	return result
}

func sortString(str string) string {
	var result []rune
	for _, r := range str {
		result = append(result, r)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return string(result)
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
