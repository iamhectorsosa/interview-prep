import { describe, expect, test } from "vitest";
import { lengthOfLastWord } from "./58-length-of-last-word";

describe("lengthOfLastWord", () => {
  test.each([
    {
      s: "",
      want: 0,
    },
    {
      s: "Hello World",
      want: 5,
    },
    {
      s: "   fly me   to   the moon  ",
      want: 4,
    },
    {
      s: "luffy is still joyboy",
      want: 6,
    },
  ])("given %s we expect %want", ({ s, want }) => {
    expect(lengthOfLastWord(s)).toBe(want);
  });
});
