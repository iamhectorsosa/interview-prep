import { describe, expect, test } from "vitest";
import { climbStairs } from "./70-climbing-stars";

describe("sqrtx", () => {
  test.each([
    {
      n: 2,
      want: 2,
    },
    {
      n: 3,
      want: 3,
    },
  ])("for $n stairs there should be $want ways to the stop", ({ n, want }) => {
    expect(climbStairs(n)).toBe(want);
  });
});
