from collections import deque


class HitCounter:
    def __init__(self):
        self.hits = deque()

    def hit(self, timestamp: int):
        self.hits.append(timestamp)

    def getHits(self, timestamp: int) -> int:
        self.clean_old_hits(timestamp)
        return len(self.hits)

    def clean_old_hits(self, current_time: int):
        threshold = current_time - 5 * 60
        while self.hits and self.hits[0] <= threshold:
            self.hits.popleft()


def solve():
    counter = HitCounter()
    counter.hit(1)
    counter.hit(2)
    counter.hit(3)
    print(counter.getHits(4))
    counter.hit(300)
    print(counter.getHits(300))
    print(counter.getHits(301))
