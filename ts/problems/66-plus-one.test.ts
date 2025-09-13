import { describe, expect, test } from "vitest";
import { plusOne } from "./66-plus-one";

describe("plusOne", () => {
  test.each([
    {
      digits: [0],
      want: [1],
    },
    {
      digits: [1, 2, 3],
      want: [1, 2, 4],
    },
    {
      digits: [4, 3, 2, 1],
      want: [4, 3, 2, 2],
    },
    {
      digits: [9],
      want: [1, 0],
    },
  ])("given $digits we want $want", ({ digits, want }) => {
    expect(plusOne(digits)).toStrictEqual(want);
  });
});
