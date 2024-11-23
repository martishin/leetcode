interface Pair {
  num: number;
  idx: number;
}

function twoSum(nums: number[], target: number): number[] {
  const numsWithIndices: Pair[] = nums
    .map((num, idx) => ({ num, idx }))
    .sort((a, b) => a.num - b.num);

  for (let idx = 0; idx < numsWithIndices.length; idx++) {
    const { num, idx: originalIdx } = numsWithIndices[idx];
    const find = target - num;

    let l = idx + 1;
    let r = numsWithIndices.length;

    while (l < r) {
      const m = Math.floor((l + r) / 2);

      if (numsWithIndices[m].num === find) {
        return [originalIdx, numsWithIndices[m].idx];
      } else if (numsWithIndices[m].num < find) {
        l = m + 1;
      } else {
        r = m;
      }
    }
  }

  return [-1, -1];
}

export function test() {
  console.log(twoSum([3, 2, 4], 6));
}
