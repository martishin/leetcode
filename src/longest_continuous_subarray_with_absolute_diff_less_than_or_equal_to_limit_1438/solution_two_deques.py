from collections import deque
from typing import List


class Solution:
    def longest_subarray(self, nums: List[int], limit: int) -> int:
        min_deque = deque()
        max_deque = deque()

        left = 0
        longest_subarray = 0

        for right, num in enumerate(nums):
            while min_deque and min_deque[-1] > num:
                min_deque.pop()
            min_deque.append(num)

            while max_deque and max_deque[-1] < num:
                max_deque.pop()
            max_deque.append(num)

            while max_deque[0] - min_deque[0] > limit:
                if max_deque[0] == nums[left]:
                    max_deque.popleft()
                if min_deque[0] == nums[left]:
                    min_deque.popleft()
                left += 1
            longest_subarray = max(longest_subarray, right - left + 1)

        return longest_subarray


def test():
    obj = Solution()
    print(obj.longest_subarray([8, 2, 4, 7], 4))
    print(obj.longest_subarray([10, 1, 2, 4, 7, 2], 5))
    print(obj.longest_subarray([4, 2, 2, 2, 4, 4, 2, 2], 0))
