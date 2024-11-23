function twoSum(nums: number[], target: number): number[] {
  const lookupTable: { [key: number]: number } = {};

  for (let idx = 0; idx < nums.length; idx++) {
    const num = nums[idx];
    const find = target - num;

    let idx_find = lookupTable[find];
    if (idx_find !== undefined) {
      return [idx_find, idx];
    }

    lookupTable[num] = idx;
  }

  return [-1, -1];
}

export function test() {
  console.log(twoSum([3, 2, 4], 6));
}
