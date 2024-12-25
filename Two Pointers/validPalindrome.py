# Given a string s, return true if it is a palindrome, otherwise return false.
# A palindrome is a word that reads the same backward as forward. For example, "racecar" is a palindrome.
# leetcode.com/problems/valid-palindrome

from typing import List

class Solution:
	def isPalindrome(self, s: str) -> bool:
		s = ''.join(e for e in s if e.isalnum()).lower()
		return s == s[::-1]

def main():
	s = "Was it a car or a cat I saw?"
	solution = Solution()
	result = solution.isPalindrome(s)
	print(result)

if __name__ == '__main__':
	main()
