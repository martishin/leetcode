function maxRob(nums: number[], cache: Map<number, number>, idx: number): number {
  if (idx >= nums.length) {
    return 0;
  }

  if (cache.has(idx)) {
    return cache.get(idx)!;
  }

  const result = Math.max(nums[idx] + maxRob(nums, cache, idx + 2), maxRob(nums, cache, idx + 1));

  cache.set(idx, result);
  return result;
}

function rob(nums: number[]): number {
  const cache = new Map<number, number>();
  return maxRob(nums, cache, 0);
}

export function test(): void {
  console.log(rob([1, 2, 3, 1]));
  console.log(rob([2, 7, 9, 3, 1]));
}
