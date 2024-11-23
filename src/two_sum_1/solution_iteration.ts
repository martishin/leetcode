interface Pair {
  idx: number;
  num: number;
}

function twoSum(nums: number[], target: number): number[] {
  const numsWithIndices: Pair[] = nums
    .map((num, idx) => ({ idx, num }))
    .sort((lhs, rhs) => lhs.num - rhs.num);

  let l = 0;
  let r = numsWithIndices.length - 1;

  while (l < r) {
    const sum = numsWithIndices[l].num + numsWithIndices[r].num;
    if (sum < target) {
      l++;
    } else if (sum > target) {
      r--;
    } else {
      return [numsWithIndices[l].idx, numsWithIndices[r].idx];
    }
  }

  return [-1, -1];
}

export function test() {
  console.log(twoSum([3, 2, 4], 6));
}
