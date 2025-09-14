import { describe, expect, test } from "vitest";
import { addBinary } from "./67-add-binary";

describe("addBinary", () => {
  test.each([
    {
      a: "11",
      b: "1",
      want: "100",
    },
    {
      a: "11",
      b: "11",
      want: "110",
    },
    {
      a: "1010",
      b: "1011",
      want: "10101",
    },
  ])("given $a and $b binary sum we want %want", ({ a, b, want }) => {
    expect(addBinary(a, b)).toEqual(want);
  });
});
