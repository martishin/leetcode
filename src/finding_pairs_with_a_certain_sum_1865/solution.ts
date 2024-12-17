class FindSumPairs {
  private readonly nums1: number[];
  private readonly nums2: number[];
  private nums2ToCount: Map<number, number>;

  constructor(nums1: number[], nums2: number[]) {
    this.nums1 = [...nums1].sort((a, b) => a - b);
    this.nums2 = nums2;
    this.nums2ToCount = new Map();

    for (const num of nums2) {
      this.nums2ToCount.set(num, (this.nums2ToCount.get(num) || 0) + 1);
    }
  }

  add(index: number, val: number): void {
    if (index < 0 || index >= this.nums2.length) {
      throw new Error("Index out of bounds");
    }

    const oldValue = this.nums2ToCount.get(this.nums2[index]);
    if (oldValue !== undefined) {
      this.nums2ToCount.set(this.nums2[index], oldValue - 1);
    }
    this.nums2[index] += val;
    this.nums2ToCount.set(this.nums2[index], (this.nums2ToCount.get(this.nums2[index]) || 0) + 1);
  }

  count(tot: number): number {
    let result = 0;

    for (const num1 of this.nums1) {
      if (num1 >= tot) {
        break;
      }

      result += this.nums2ToCount.get(tot - num1) || 0;
    }

    return result;
  }
}

export function test(): void {
  const obj = new FindSumPairs([1, 1, 2, 2, 2, 3], [1, 4, 5, 2, 5, 4]);
  console.log(obj.count(7));
  obj.add(3, 2);
  console.log(obj.count(8));
  console.log(obj.count(4));
  obj.add(0, 1);
  obj.add(1, 1);
  console.log(obj.count(7));
}
