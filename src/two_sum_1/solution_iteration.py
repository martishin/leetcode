from typing import List


class Solution:
    def two_sum(self, nums: List[int], target: int) -> List[int]:
        nums_with_indices = sorted((num, idx) for num, idx in enumerate(nums))

        l = 0
        r = len(nums_with_indices) - 1

        while l < r:
            nums_sum = nums_with_indices[l][0] + nums_with_indices[r][0]
            if nums_sum < target:
                l += 1
            elif nums_sum > target:
                r -= 1
            else:
                return [nums_with_indices[l][1], nums_with_indices[r][1]]

        return [-1, -1]


def test():
    print(Solution().two_sum([3, 2, 4], 6))
    print(Solution().two_sum([0, 4, 3, 0], 0))
