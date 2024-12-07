# Given two strings s and t, return true if the two strings are anagrams of each other, otherwise return false.
# An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.

class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        def get_count_map(s: str) -> dict:
            res = {}
            for char in s:
                if char not in res:
                    res[char] = 0
                res[char] += 1
            return res

        scm = get_count_map(s)
        tcm = get_count_map(t)

        for key, val in scm.items():
            if scm.get(key, 0) != tcm.get(key, 0):
                return False

        for key in tcm:
            if key not in scm:
                return False

        return True


def main():
    solution = Solution()
    s = "racecar"
    t = "carrace"
    print(f"Test case 1: {solution.isAnagram(s, t)}")

    s = "jar"
    t = "jam"
    print(f"Test case 2: {solution.isAnagram(s, t)}")

if __name__ == "__main__":
    main()
