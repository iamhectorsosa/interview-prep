import { describe, expect, test } from "vitest";
import { removeDuplicates } from "./26-remove-duplicates-from-sorted-array";

describe("removeDuplicates", () => {
  test.each([
    {
      nums: [],
      want: 0,
    },
    {
      nums: [1, 1, 2],
      want: 2,
    },
    {
      nums: [0, 0, 1, 1, 1, 2, 2, 3, 3, 4],
      want: 5,
    },
  ])("given $nums, we want $want", ({ nums, want }) => {
    expect(removeDuplicates(nums)).toStrictEqual(want);
  });
});
