# Given an integer array nums, return true if any value appears more than once in the array, otherwise return false.

from typing import List

class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        res = {}
        for val in nums:
            if val in res:
                return True
            res[val] = val
        return False

def main():
    solution = Solution()
    nums1 = [1, 2, 3, 4, 5]
    print(f"Test case 1: {solution.hasDuplicate(nums1)}")

    nums2 = [1, 2, 3, 3, 4]
    print(f"Test case 2: {solution.hasDuplicate(nums2)}")

if __name__ == "__main__":
    main()
