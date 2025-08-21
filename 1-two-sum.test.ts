import { describe, expect, test } from "vitest";
import { twoSum } from "./1-two-sum";

describe("twoSum", () => {
  test.each([
    {
      nums: [],
      target: 0,
      expected: [],
    },
    {
      nums: [2, 7, 11, 15],
      target: 9,
      expected: [0, 1],
    },
    {
      nums: [3, 2, 4],
      target: 6,
      expected: [1, 2],
    },
    {
      nums: [3, 3],
      target: 6,
      expected: [0, 1],
    },
  ])(
    "from $nums with target of $target we want $expected",
    ({ nums, target, expected }) => {
      expect(twoSum(nums, target)).toStrictEqual(expected);
    },
  );
});
