function rob(nums: number[]): number {
  const n = nums.length;

  if (n === 1) {
    return nums[0];
  }

  if (n === 2) {
    return Math.max(nums[0], nums[1]);
  }

  return Math.max(robHelper(nums.slice(1)), robHelper(nums.slice(0, n - 1)));
}

function robHelper(nums: number[]): number {
  let secondLast = nums[0];
  let last = Math.max(nums[0], nums[1]);

  for (let i = 2; i < nums.length; i++) {
    const current = Math.max(last, secondLast + nums[i]);
    secondLast = last;
    last = current;
  }

  return last;
}

export function test(): void {
  console.log(rob([1, 2, 3, 1]));
  console.log(rob([2, 7, 9, 3, 1]));
}
