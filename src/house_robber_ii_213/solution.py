from typing import List


class Solution:
    # def rob(self, nums: List[int]) -> int:
    #     n = len(nums)
    #
    #     if n == 0:
    #         return 0
    #     elif n == 1:
    #         return nums[0]
    #
    #     dp = [0] * n
    #     dp[0] = nums[0]
    #     dp[1] = max(nums[0], nums[1])
    #
    #     for i in range(2, n):
    #         dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])
    #
    #     return dp[n - 1]

    def rob(self, nums: List[int]) -> int:
        n = len(nums)

        if n == 0:
            return 0
        elif n == 1:
            return nums[0]

        prev2 = nums[0]
        prev1 = max(nums[0], nums[1])

        for i in range(2, n):
            current = max(prev1, prev2 + nums[i])
            prev2 = prev1
            prev1 = current

        return prev1


def test():
    solution = Solution()
    print(solution.rob([1, 2, 3, 1]))
    print(solution.rob([2, 7, 9, 3, 1]))
