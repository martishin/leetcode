from collections import defaultdict
from typing import Tuple, List, Dict


class TimeMap:
    def __init__(self):
        self._key_to_values: Dict[str, List[Tuple[int, str]]] = defaultdict(list)

    def set(self, key: str, value: str, timestamp: int) -> None:
        self._key_to_values[key].append((timestamp, value))

    def get(self, key: str, timestamp: int) -> str:
        values = self._key_to_values.get(key)
        if values is None:
            return ""

        left = 0
        right = len(values)

        while left < right:
            mid = (left + right) // 2
            if values[mid][0] <= timestamp:
                left = mid + 1
            else:
                right = mid

        if right == 0:
            return ""

        return values[right - 1][1]


# Your TimeMap object will be instantiated and called as such:
# obj = TimeMap()
# obj.set(key,value,timestamp)
# param_2 = obj.get(key,timestamp)
def test():
    obj = TimeMap()
    obj.set("foo", "bar", 1)
    print(obj.get("foo", 1))
    print(obj.get("foo", 3))
    obj.set("foo", "bar2", 4)
    print(obj.get("foo", 4))
    print(obj.get("foo", 5))
