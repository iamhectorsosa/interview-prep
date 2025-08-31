import { describe, expect, test } from "vitest";
import { strStr } from "./28-find-the-index-of-the-first-occurrence-in-a-string";

describe("strStr", () => {
  test.each([
    {
      haystack: "",
      needle: "",
      want: 0,
    },
    {
      haystack: "sadbutsad",
      needle: "sad",
      want: 0,
    },
    {
      haystack: "leetcode",
      needle: "leeto",
      want: -1,
    },
    {
      haystack: "hello",
      needle: "ll",
      want: 2,
    },
    {
      haystack: "mississippi",
      needle: "issi",
      want: 1,
    },
    {
      haystack: "hector",
      needle: "or",
      want: 4,
    },
  ])(
    "if we put $needle in $haystack its first occurrance index is at $want",
    ({ haystack, needle, want }) => {
      expect(strStr(haystack, needle)).toBe(want);
    },
  );
});
