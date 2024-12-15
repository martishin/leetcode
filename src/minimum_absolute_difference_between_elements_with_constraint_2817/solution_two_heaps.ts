class MinHeap<T> {
  private heap: T[] = [];

  push(item: T): void {
    this.heap.push(item);
    this.bubbleUp();
  }

  pop(): T | undefined {
    if (this.size() === 0) return undefined;
    this.swap(0, this.heap.length - 1);
    const item = this.heap.pop();
    this.bubbleDown();
    return item;
  }

  peek(): T | undefined {
    return this.heap[0];
  }

  size(): number {
    return this.heap.length;
  }

  private bubbleUp(): void {
    let index = this.heap.length - 1;
    const current = this.heap[index];

    while (index > 0) {
      const parentIndex = Math.floor((index - 1) / 2);
      const parent = this.heap[parentIndex];
      if (this.compare(current, parent) >= 0) break;

      this.swap(index, parentIndex);
      index = parentIndex;
    }
  }

  private bubbleDown(): void {
    let index = 0;

    while (true) {
      const leftIndex = 2 * index + 1;
      const rightIndex = 2 * index + 2;
      let smallestIndex = index;

      if (
        leftIndex < this.heap.length &&
        this.compare(this.heap[leftIndex], this.heap[smallestIndex]) < 0
      ) {
        smallestIndex = leftIndex;
      }
      if (
        rightIndex < this.heap.length &&
        this.compare(this.heap[rightIndex], this.heap[smallestIndex]) < 0
      ) {
        smallestIndex = rightIndex;
      }
      if (smallestIndex === index) break;

      this.swap(index, smallestIndex);
      index = smallestIndex;
    }
  }

  private swap(i: number, j: number): void {
    [this.heap[i], this.heap[j]] = [this.heap[j], this.heap[i]];
  }

  protected compare(a: T, b: T): number {
    return (a as any) - (b as any); // Override this for custom comparisons
  }
}

class MinHeapWithCustomComparator<T> extends MinHeap<[number, T]> {
  protected compare(a: [number, T], b: [number, T]): number {
    return a[0] - b[0];
  }
}

class MaxHeapWithCustomComparator<T> extends MinHeap<[number, T]> {
  protected compare(a: [number, T], b: [number, T]): number {
    return b[0] - a[0];
  }
}

function minDiff(nums: number[], x: number): number {
  const sortedList: [number, number][] = nums
    .map<[number, number]>((val, idx) => [val, idx])
    .sort((a, b) => a[0] - b[0]);

  const minHeap = new MinHeapWithCustomComparator<number>();
  const maxHeap = new MaxHeapWithCustomComparator<number>();
  let result = Infinity;

  for (const [val, idx] of sortedList) {
    minHeap.push([idx, val]);
    maxHeap.push([idx, val]);

    while (minHeap.size() > 0 && minHeap.peek()![0] <= idx - x) {
      result = Math.min(result, Math.abs(val - minHeap.pop()![1]));
    }

    while (maxHeap.size() > 0 && maxHeap.peek()![0] >= idx + x) {
      result = Math.min(result, Math.abs(val - maxHeap.pop()![1]));
    }
  }

  return result;
}

export function test() {
  console.log(minDiff([4, 3, 2, 4], 2));
  console.log(minDiff([5, 3, 2, 10, 15], 1));
  console.log(minDiff([1, 2, 3, 4], 3));
  console.log(minDiff([14, 111, 16], 1));
}
