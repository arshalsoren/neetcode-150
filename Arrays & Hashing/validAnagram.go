// Given two strings s and t, return true if the two strings are anagrams of each other, otherwise return false.
// An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.

package main

import "fmt"

type Solution struct{}

func (s *Solution) isAnagram(str1 string, str2 string) bool {
	getCountMap := func(s string) map[rune]int {
		res := make(map[rune]int)
		for _, char := range s {
			res[char]++
		}
		return res
	}

	scm := getCountMap(str1)
	tcm := getCountMap(str2)

	for key := range scm {
		if scm[key] != tcm[key] {
			return false
		}
	}

	for key := range tcm {
		if _, exists := scm[key]; !exists {
			return false
		}
	}

	return true
}

func main() {
	solution := &Solution{}
	s := "racecar"
	t := "carrace"
	fmt.Printf("Test case 1: %v\n", solution.isAnagram(s, t))

	s = "jar"
	t = "jam"
	fmt.Printf("Test case 2: %v\n", solution.isAnagram(s, t))
}
