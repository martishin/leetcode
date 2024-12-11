from typing import List

from sortedcontainers import SortedList


def min_diff(nums: List[int], x: int) -> int:
    sorted_list = SortedList()

    result = float("inf")
    for i in range(x, len(nums)):
        sorted_list.add(nums[i - x])

        upper_bound = sorted_list.bisect_right(nums[i])

        if upper_bound < len(sorted_list):
            result = min(result, abs(nums[i] - sorted_list[upper_bound]))
        if upper_bound > 0:
            result = min(result, abs(nums[i] - sorted_list[upper_bound - 1]))

    return result


def test():
    print(min_diff([4, 3, 2, 4], 2))
    print(min_diff([5, 3, 2, 10, 15], 1))
    print(min_diff([1, 2, 3, 4], 3))
    print(min_diff([14, 111, 16], 1))
