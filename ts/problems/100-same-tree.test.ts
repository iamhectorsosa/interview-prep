import { describe, expect, test } from "vitest";
import { isSameTree } from "./100-same-tree";

describe("isSameTree", () => {
  test.each([
    {
      p: {
        val: 1,
        left: {
          val: 2,
          left: null,
          right: null,
        },
        right: {
          val: 3,
          left: null,
          right: null,
        },
      },
      q: {
        val: 1,
        left: {
          val: 2,
          left: null,
          right: null,
        },
        right: {
          val: 3,
          left: null,
          right: null,
        },
      },
      want: true,
    },
  ])("given trees $p and $q", ({ p, q, want }) => {
    expect(isSameTree(p, q)).toBe(want);
  });
});
