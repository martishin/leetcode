from typing import List


class Solution:
    def two_sum(self, nums: List[int], target: int) -> List[int]:
        nums_with_indices = sorted((num, idx) for idx, num in enumerate(nums))

        for idx, (num, original_idx) in enumerate(nums_with_indices):
            l = idx + 1
            r = len(nums_with_indices)

            while l < r:
                m = (l + r) // 2
                find = target - num

                if nums_with_indices[m][0] == find:
                    return [original_idx, nums_with_indices[m][1]]
                elif nums_with_indices[m][0] < find:
                    l = m + 1
                else:
                    r = m

        return [-1, -1]


def test():
    print(Solution().two_sum([3, 2, 4], 6))
    print(Solution().two_sum([0, 4, 3, 0], 0))
