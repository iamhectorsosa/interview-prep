export function removeDuplicates(nums: Array<number>): number {
  if (!nums.length || nums.length === 1) return nums.length;

  // nums.splice(0, nums.length, ...new Set(nums));

  const map: Record<number, string> = {};
  for (let i = nums.length - 1; i >= 0; i--) {
    if (nums[i] in map) {
      nums.splice(i, 1);
      continue;
    }
    map[nums[i]] = "";
  }

  return nums.length;
}
