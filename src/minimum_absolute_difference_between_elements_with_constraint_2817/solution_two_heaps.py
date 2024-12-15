import heapq
from typing import List


def min_diff(nums: List[int], x: int) -> int:
    sorted_list = sorted((val, idx) for idx, val in enumerate(nums))

    min_heap = []
    max_heap = []
    result = 1e10

    for val, idx in sorted_list:
        heapq.heappush(min_heap, (idx, val))
        heapq.heappush(max_heap, (-idx, val))

        while len(min_heap) > 0 and min_heap[0][0] <= idx - x:
            result = min(result, val - heapq.heappop(min_heap)[1])
        while len(max_heap) > 0 and max_heap[0][0] <= -idx - x:
            result = min(result, val - heapq.heappop(max_heap)[1])

    return result


def test():
    print(min_diff([4, 3, 2, 4], 2))
    print(min_diff([5, 3, 2, 10, 15], 1))
    print(min_diff([1, 2, 3, 4], 3))
    print(min_diff([14, 111, 16], 1))
