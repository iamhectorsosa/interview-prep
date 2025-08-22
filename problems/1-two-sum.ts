export function twoSum(nums: Array<number>, target: number): Array<number> {
  const h: Record<number, number> = {};
  for (let i = 0; i < nums.length; i++) {
    const d = target - nums[i];
    if (Object.hasOwn(h, d)) {
      return [h[d], i];
    }
    h[nums[i]] = i;
  }
  return [];
}
