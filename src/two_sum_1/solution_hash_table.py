from typing import List


class Solution:
    def two_sum(self, nums: List[int], target: int) -> List[int]:
        lookup_table = {}

        for idx, num in enumerate(nums):
            find = target - num

            idx_find = lookup_table.get(find)
            if idx_find is not None:
                return [idx_find, idx]

            lookup_table[num] = idx

        return [-1, -1]


def test():
    print(Solution().two_sum([3, 2, 4], 6))
    print(Solution().two_sum([0, 4, 3, 0], 0))
