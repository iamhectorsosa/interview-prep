import { describe, expect, test } from "vitest";
import { twoSum } from "./1-two-sum";

describe("twoSum", () => {
  test.each([
    {
      nums: [],
      target: 0,
      want: [],
    },
    {
      nums: [2, 7, 11, 15],
      target: 9,
      want: [0, 1],
    },
    {
      nums: [3, 2, 4],
      target: 6,
      want: [1, 2],
    },
    {
      nums: [3, 3],
      target: 6,
      want: [0, 1],
    },
  ])(
    "from $nums with target of $target we want $want",
    ({ nums, target, want }) => {
      expect(twoSum(nums, target)).toStrictEqual(want);
    },
  );
});
