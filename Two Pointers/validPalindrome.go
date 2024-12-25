// Given a string s, return true if it is a palindrome, otherwise return false.

package main

import "fmt"

type Solution struct{}

func (s *Solution) isPalindrome(str string) bool {
	l, r := 0, len(str)-1
	for l < r {
		if !isAlphanumeric(str[l]) {
			l++
		} else if !isAlphanumeric(str[r]) {
			r--
		} else if toLower(str[l]) != toLower(str[r]) {
			return false
		} else {
			l++
			r--
		}
	}
	return true
}

func isAlphanumeric(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}

func toLower(ch byte) byte {
	if ch >= 'A' && ch <= 'Z' {
		return ch + 32
	}
	return ch
}

func main() {
	solution := &Solution{}
	s1 := "Was it a car or a cat I saw?"
	fmt.Printf("Test case 1: %v\n", solution.isPalindrome(s1))

	s2 := "tab a cat"
	fmt.Printf("Test case 2: %v\n", solution.isPalindrome(s2))
}
