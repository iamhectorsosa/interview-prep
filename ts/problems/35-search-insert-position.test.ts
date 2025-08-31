import { describe, expect, test } from "vitest";
import { searchInsert } from "./35-search-insert-position";

describe("searchInsert", () => {
  test.each([
    {
      nums: [],
      target: 0,
      want: 0,
    },
    {
      nums: [1, 3, 5, 6],
      target: 5,
      want: 2,
    },
    {
      nums: [1, 3, 5, 6],
      target: 2,
      want: 1,
    },
    {
      nums: [1, 3, 5, 6],
      target: 7,
      want: 4,
    },
  ])(
    "for $nums with a target of $target, the search position should be $want",
    ({ nums, target, want }) => {
      expect(searchInsert(nums, target)).toBe(want);
    },
  );
});
