export function searchInsert(nums: number[], target: number): number {
  if (!nums.length) return 0;
  // for (let i = 0; i < nums.length; i++) {
  //   if (nums[i] === target || nums[i] > target) return i;
  // }
  // return nums.length;

  const existingIndex = nums.findIndex((el) => el == target);
  if (existingIndex != -1) {
    return existingIndex;
  }

  let leftBoundary = 0;
  let rightBoundary = nums.length - 1;

  while (leftBoundary <= rightBoundary) {
    let middle = Math.floor((leftBoundary + rightBoundary) / 2);
    if (target > nums[middle]) {
      leftBoundary = middle + 1;
    } else {
      rightBoundary = middle - 1;
    }
  }

  return leftBoundary;
}
