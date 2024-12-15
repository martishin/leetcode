from collections import Counter
from typing import List


class FindSumPairs:
    def __init__(self, nums1: List[int], nums2: List[int]):
        self._nums1 = sorted(nums1)
        self._nums2 = nums2
        self._nums2ToCount = Counter(nums2)

    def add(self, index: int, val: int) -> None:
        self._nums2ToCount[self._nums2[index]] -= 1
        self._nums2[index] += val
        self._nums2ToCount[self._nums2[index]] += 1

    def count(self, tot: int) -> int:
        result = 0

        for num1 in self._nums1:
            if num1 >= tot:
                break

            result += self._nums2ToCount[tot - num1]

        return result


def test():
    obj = FindSumPairs([1, 1, 2, 2, 2, 3], [1, 4, 5, 2, 5, 4])
    print(obj.count(7))
    obj.add(3, 2)
    print(obj.count(8))
    print(obj.count(4))
    obj.add(0, 1)
    obj.add(1, 1)
    print(obj.count(7))
