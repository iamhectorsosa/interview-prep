import { describe, expect, test } from "vitest";
import { isValid } from "./20-valid-parentheses";

describe("isValid", () => {
  test.each([
    {
      s: "()",
      want: true,
    },
    {
      s: "()[]{}",
      want: true,
    },
    {
      s: "(]",
      want: false,
    },
    {
      s: "([])",
      want: true,
    },
    {
      s: "([)]",
      want: false,
    },
    {
      s: "([]",
      want: false,
    },
  ])("$s expected to have valid parentheses is $want", ({ s, want }) => {
    expect(isValid(s)).toBe(want);
  });
});
