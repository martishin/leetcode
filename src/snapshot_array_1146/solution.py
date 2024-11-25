class Snapshot:
    def __init__(self, version: int, val: int):
        self.version = version
        self.val = val


class SnapshotArray:
    def __init__(self, length: int):
        self._entries = [[] for _ in range(length)]
        self._version = 0

    def set(self, index: int, val: int) -> None:
        if self._entries[index]:
            last_entry = self._entries[index][-1]
            if last_entry.version == self._version:
                last_entry.val = val
                return

        self._entries[index].append(Snapshot(self._version, val))

    def snap(self) -> int:
        self._version += 1
        return self._version - 1

    def get(self, index: int, snap_id: int) -> int:
        if not self._entries[index]:
            return 0

        idx_entries = self._entries[index]

        l, r = 0, len(idx_entries)
        while l < r:
            m = (l + r) // 2
            if idx_entries[m].version == snap_id:
                return idx_entries[m].val
            elif idx_entries[m].version < snap_id:
                l = m + 1
            else:
                r = m

        if r == 0:
            return 0

        return idx_entries[r - 1].val


# Your SnapshotArray object will be instantiated and called as such:
# obj = SnapshotArray(length)
# obj.set(index,val)
# param_2 = obj.snap()
# param_3 = obj.get(index,snap_id)
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
