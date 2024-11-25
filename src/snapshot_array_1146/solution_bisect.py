import bisect


class SnapshotArray:
    def __init__(self, length: int):
        self._entries = [[[0, 0]] for _ in range(length)]
        self._version = 0

    def set(self, index: int, val: int) -> None:
        if self._entries[index][-1][0] == self._version:
            self._entries[index][-1][1] = val
        else:
            self._entries[index].append([self._version, val])

    def snap(self) -> int:
        self._version += 1
        return self._version - 1

    def get(self, index: int, snap_id: int) -> int:
        snap_index = bisect.bisect_right(self._entries[index], [snap_id, float("inf")])
        return self._entries[index][snap_index - 1][1]


def test():
    obj = SnapshotArray(3)
    obj.set(0, 5)
    print(obj.snap(), end=",")
    obj.set(0, 6)
    print(obj.get(0, 0))

    obj = SnapshotArray(2)
    print(obj.snap(), end=",")
    print(obj.get(1, 0), end=",")
    print(obj.get(0, 0), end=",")
    obj.set(1, 8)
    print(obj.get(1, 0), end=",")
    obj.set(0, 20)
    print(obj.get(0, 0))
    obj.set(0, 7)
