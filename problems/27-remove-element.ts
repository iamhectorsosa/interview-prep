export function removeElement(nums: number[], val: number): number {
  if (!nums.length || (nums.length === 1 && nums[0] !== val)) {
    return nums.length;
  }

  for (let i = nums.length; i >= 0; i--) {
    if (nums[i] === val) {
      nums.splice(i, 1);
    }
  }

  return nums.length;
}
