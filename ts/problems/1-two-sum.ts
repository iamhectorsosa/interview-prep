export function twoSum(nums: Array<number>, target: number): Array<number> {
  const hashMap: Record<number, number> = {};
  for (let i = 0; i < nums.length; i++) {
    const diff = target - nums[i];
    if (Object.hasOwn(hashMap, diff)) {
      return [hashMap[diff], i];
    }
    hashMap[nums[i]] = i;
  }
  return [];
}
