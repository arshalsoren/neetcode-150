# Given an integer array nums, return an array output where output[i] is the product of all the elements of nums except nums[i].
# Each product is guaranteed to fit in a 32-bit integer.
# https://leetcode.com/problems/products-of-array-except-self/

class Solution:
    def productExceptSelfV2(self, nums):
        res = [1] * len(nums)
        prefix = 1
        for i in range(len(nums)):
            res[i] = prefix
            prefix *= nums[i]
        postfix = 1
        for i in range(len(nums) - 1, -1, -1):
            res[i] *= postfix
            postfix *= nums[i]
        return res

    def productExceptSelf(self, nums):
        output = []
        for i in range(len(nums)):
            product = 1;
            for j in range(len(nums)):
                if (i == j):
                    continue
                product *= nums[j]

            output.append(product)
        return output


def main():
    nums = [1, 2, 4, 6]
    print(Solution().productExceptSelf(nums))
    print(Solution().productExceptSelfV2(nums))

if __name__ == "__main__":
    main()
