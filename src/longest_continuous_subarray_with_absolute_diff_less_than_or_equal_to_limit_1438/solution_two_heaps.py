import heapq
from typing import List


class Solution:
    def longest_subarray(self, nums: List[int], limit: int) -> int:
        min_heap = []
        max_heap = []

        left = 0
        longest_subarray = 0

        for right, num in enumerate(nums):
            heapq.heappush(min_heap, (num, right))
            heapq.heappush(max_heap, (-num, right))

            while -max_heap[0][0] - min_heap[0][0] > limit:
                left = min(max_heap[0][1], min_heap[0][1]) + 1

                while max_heap[0][1] < left:
                    heapq.heappop(max_heap)

                while min_heap[0][1] < left:
                    heapq.heappop(min_heap)

            longest_subarray = max(longest_subarray, right - left + 1)

        return longest_subarray


def test():
    obj = Solution()
    print(obj.longest_subarray([8, 2, 4, 7], 4))
    print(obj.longest_subarray([10, 1, 2, 4, 7, 2], 5))
    print(obj.longest_subarray([4, 2, 2, 2, 4, 4, 2, 2], 0))
