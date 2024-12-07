# Given an array of integers nums and an integer target, return the indices i and j such that nums[i] + nums[j] == target and i != j.
# You may assume that every input has exactly one pair of indices i and j that satisfy the condition.
# Return the answer with the smaller index first.
# https://leetcode.com/problems/two-sum/
from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        diff_pos_map = {}
        for pos, val in enumerate(nums):
            diff = target - val
            if val in diff_pos_map:
                return [diff_pos_map[val], pos]
            diff_pos_map[diff] = pos

def main():
    solution = Solution()
    nums = [3,4,5,6]
    target = 7
    print(f"Test case 1: {solution.twoSum(nums, target)}")
    nums = [4,5,6]
    target = 10
    print(f"Test case 2: {solution.twoSum(nums, target)}")
    nums = [5,5]
    target = 10
    print(f"Test case 3: {solution.twoSum(nums, target)}")

if __name__ == "__main__":
    main()
