from typing import List

from sortedcontainers import SortedList


class Solution:
    def longest_subarray(self, nums: List[int], limit: int) -> int:
        sorted_list = SortedList()

        left = 0
        longest_subarray = 0

        for right, num in enumerate(nums):
            sorted_list.add(num)

            while sorted_list[-1] - sorted_list[0] > limit:
                sorted_list.remove(nums[left])
                left += 1

            longest_subarray = max(longest_subarray, len(sorted_list))

        return longest_subarray


def test():
    obj = Solution()
    print(obj.longest_subarray([8, 2, 4, 7], 4))
    print(obj.longest_subarray([10, 1, 2, 4, 7, 2], 5))
    print(obj.longest_subarray([4, 2, 2, 2, 4, 4, 2, 2], 0))
