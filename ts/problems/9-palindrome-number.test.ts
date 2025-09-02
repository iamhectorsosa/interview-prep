import { describe, expect, test } from "vitest";
import { isPalindrome, isPalindromeStr } from "./9-palindrome-number";

describe("isPalindrome", () => {
  test.each([
    {
      x: 121,
      want: true,
    },
    {
      x: -121,
      want: false,
    },
    {
      x: 10,
      want: false,
    },
    {
      x: 0,
      want: true,
    },
  ])("check if $x is a palindrome and expects $want", ({ x, want }) => {
    expect(isPalindrome(x)).toBe(want);
  });
});

describe("isPalindromeStr", () => {
  test.each([
    {
      x: 121,
      want: true,
    },
    {
      x: -121,
      want: false,
    },
    {
      x: 10,
      want: false,
    },
    {
      x: 0,
      want: true,
    },
  ])("check if $x is a palindrome and expects $want", ({ x, want }) => {
    expect(isPalindromeStr(x)).toBe(want);
  });
});
