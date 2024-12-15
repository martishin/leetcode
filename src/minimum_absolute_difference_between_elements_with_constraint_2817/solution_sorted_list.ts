class SortedList {
  private list: number[];

  constructor() {
    this.list = [];
  }

  add(value: number): void {
    const index = this.bisectRight(value);
    this.list.splice(index, 0, value);
  }

  bisectRight(value: number): number {
    let low = 0,
      high = this.list.length;
    while (low < high) {
      const mid = Math.floor((low + high) / 2);
      if (this.list[mid] < value) {
        low = mid + 1;
      } else {
        high = mid;
      }
    }
    return high;
  }

  get(index: number): number {
    return this.list[index];
  }

  length(): number {
    return this.list.length;
  }
}

function minDiff(nums: number[], x: number): number {
  const sortedList = new SortedList();
  let result = Infinity;

  for (let i = x; i < nums.length; i++) {
    sortedList.add(nums[i - x]);

    const upperBound = sortedList.bisectRight(nums[i]);

    if (upperBound < sortedList.length()) {
      result = Math.min(result, Math.abs(nums[i] - sortedList.get(upperBound)));
    }
    if (upperBound > 0) {
      result = Math.min(result, Math.abs(nums[i] - sortedList.get(upperBound - 1)));
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
