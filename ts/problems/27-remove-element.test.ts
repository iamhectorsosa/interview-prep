import { describe, expect, test } from "vitest";
import { removeElement } from "./27-remove-element";

describe("removeElement", () => {
  test.each([
    {
      nums: [],
      val: 0,
      want: 0,
    },
    {
      nums: [3, 2, 2, 3],
      val: 3,
      want: 2,
    },
    {
      nums: [0, 1, 2, 2, 3, 0, 4, 2],
      val: 2,
      want: 5,
    },
  ])("given $nums and val $val, we expect $want", ({ nums, val, want }) => {
    expect(removeElement(nums, val)).toBe(want);
  });
});
