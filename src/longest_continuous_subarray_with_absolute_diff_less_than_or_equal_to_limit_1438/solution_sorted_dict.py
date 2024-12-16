from typing import List

from sortedcontainers import SortedDict


class Solution:
    def longest_subarray(self, nums: List[int], limit: int) -> int:
        window = SortedDict()
        left = 0
        max_length = 0

        for right, num in enumerate(nums):
            window[num] = window.get(num, 0) + 1

            while window.peekitem(-1)[0] - window.peekitem(0)[0] > limit:
                window[nums[left]] -= 1
                if window[nums[left]] == 0:
                    window.pop(nums[left])

                left += 1

            max_length = max(max_length, right - left + 1)

        return max_length


def test():
    obj = Solution()
    print(obj.longest_subarray([8, 2, 4, 7], 4))
    print(obj.longest_subarray([10, 1, 2, 4, 7, 2], 5))
    print(obj.longest_subarray([4, 2, 2, 2, 4, 4, 2, 2], 0))
