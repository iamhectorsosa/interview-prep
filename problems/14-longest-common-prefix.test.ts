import { describe, expect, test } from "vitest";
import { longestCommonPrefix } from "./14-longest-common-prefix";

describe("longestCommonPrefix", () => {
  test.each([
    {
      strs: ["flower", "flow", "flight"],
      want: "fl",
    },
    {
      strs: ["dog", "racecar", "car"],
      want: "",
    },
    {
      strs: [],
      want: "",
    },
    {
      strs: ["cir", "car"],
      want: "c",
    },
    {
      strs: ["baab", "bacb", "b", "cbc"],
      want: "",
    },
  ])("for the array $strs we expect the prefix $want", ({ strs, want }) => {
    expect(longestCommonPrefix(strs)).toBe(want);
  });
});
