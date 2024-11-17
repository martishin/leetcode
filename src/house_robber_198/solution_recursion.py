from typing import List, Dict


class Solution:
    def max_rob(self, nums: List[int], cache: Dict[int, int], idx: int) -> int:
        if idx >= len(nums):
            return 0

        if idx in cache:
            return cache[idx]

        result = max(
            nums[idx] + self.max_rob(nums, cache, idx + 2),
            self.max_rob(nums, cache, idx + 1),
        )

        cache[idx] = result
        return result

    def rob(self, nums: List[int]) -> int:
        cache: Dict[int, int] = {}
        return self.max_rob(nums, cache, 0)


def test():
    solution = Solution()
    print(solution.rob([1, 2, 3, 1]))
    print(solution.rob([2, 7, 9, 3, 1]))
