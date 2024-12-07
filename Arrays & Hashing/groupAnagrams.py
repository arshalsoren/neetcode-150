# Given an array of strings strs, group all anagrams together into sublists. You may return the output in any order.
# An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.
# leetcode.com/problems/group-anagrams

from typing import List

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        anagram_groups = {}

        for word in strs:
            key = ''.join(sorted(word))

            if key not in anagram_groups:
                anagram_groups[key] = []
            anagram_groups[key].append(word)

        return list(anagram_groups.values())

def main():
    strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
    solution = Solution()
    result = solution.groupAnagrams(strs)
    print(result)

if __name__ == '__main__':
    main()
