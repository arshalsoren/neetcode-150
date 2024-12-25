# You are given an integer array heights where heights[i] represents the height of the ith bar.
# You may choose any two bars to form a container. Return the maximum amount of water a container can store.
# https://leetcode.com/problems/container-with-most-water/


from typing import List
class Solution:
    def maxArea(self, height: List[int]) -> int:
        if len(height) == 0:
            return 0
        left = 0
        right = len(height) - 1
        max_area = 0
        while left < right:
            max_area = max(max_area, min(height[left], height[right]) * (right - left))
            if height[left] < height[right]:
                left += 1
            else:
                right -= 1
        return max_area


def main():
    height = [1,8,6,2,5,4,8,3,7]
    # Output: 49
    solution = Solution()
    result = solution.maxArea(height)
    print(result)

if __name__ == '__main__':
    main()
