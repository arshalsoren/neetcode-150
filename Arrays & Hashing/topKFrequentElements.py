# Given an integer array nums and an integer k, return the k most frequent elements within the array.
# The test cases are generated such that the answer is always unique.
# You may return the output in any order.
# https://leetcode.com/problems/top-k-frequent-elements/

from typing import List
from collections import Counter

class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        frequency = Counter(nums)
        result = [item[0] for item in frequency.most_common(k)]

        return result

def main():
    nums = [1,1,1,2,2,3]
    k = 2
    solution = Solution()
    result = solution.topKFrequent(nums, k)
    print(result)

if __name__ == '__main__':
    main()
