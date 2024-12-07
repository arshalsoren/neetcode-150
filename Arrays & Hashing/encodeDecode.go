// Design an algorithm to encode a list of strings to a single string. The encoded string is then decoded back to the original list of strings.
// Please implement encode and decode
// https://www.hackerrank.com/challenges/encode-and-decode/problem

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func encode(strs []string) string {
	var res string
	for _, s := range strs {
		res += fmt.Sprintf("%d#%s", len(s), s)
	}
	return res
}

func decode(str string) []string {
	if str == "" {
		return []string{}
	}

	var res []string
	i := 0
	for i < len(str) {
		j := strings.Index(str[i:], "#")
		if j == -1 {
			break
		}
		j += i
		length, _ := strconv.Atoi(str[i:j])
		res = append(res, str[j+1:j+1+length])
		i = j + 1 + length
	}
	return res
}

func main() {
	strs := []string{"neet", "code", "love", "you"}
	s := encode(strs)
	fmt.Println("Encoded String: ", s)
	fmt.Println("Decoded String: ", decode(s))
}
