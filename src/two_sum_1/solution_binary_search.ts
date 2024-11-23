function twoSum(nums: number[], target: number): number[] {
  const numsWithIndices: [number, number][] = nums
    .map((num, idx): [number, number] => [num, idx])
    .sort((a, b) => a[0] - b[0]);

  for (let idx = 0; idx < numsWithIndices.length; idx++) {
    const [num, originalIdx] = numsWithIndices[idx];
    const find = target - num;

    let l = idx + 1;
    let r = numsWithIndices.length;

    while (l < r) {
      const m = Math.floor((l + r) / 2);

      if (numsWithIndices[m][0] === find) {
        return [originalIdx, numsWithIndices[m][1]];
      } else if (numsWithIndices[m][0] < find) {
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
