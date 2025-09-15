import { describe, expect, test } from "vitest";
import { mySqrt } from "./69-sqrtx";

describe("sqrtx", () => {
  test.each([
    {
      x: 4,
      want: 2,
    },
    {
      x: 8,
      want: 2,
    },
    {
      x: 16,
      want: 4,
    },
    {
      x: 18,
      want: 4,
    },
    {
      x: 0,
      want: 0,
    },
  ])(
    "the sqrt of $x rounded to the nearest integer should be $want",
    ({ x, want }) => {
      expect(mySqrt(x)).toBe(want);
    },
  );
});
