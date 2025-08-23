import { describe, expect, test } from "vitest";
import { romanToInt } from "./13-roman-to-integer";

describe("romanToInt", () => {
  test.each([
    {
      s: "III",
      want: 3,
    },
    {
      s: "LVIII",
      want: 58,
    },
    {
      s: "MCMXCIV",
      want: 1994,
    },
  ])("for roman $s we expect $want", ({ s, want }) => {
    expect(romanToInt(s)).toBe(want);
  });
});
