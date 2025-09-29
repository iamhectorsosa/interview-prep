import { describe, expect, test } from "vitest";
import { inorderTraversal } from "./94-binary-tree-inorder-traversal";

describe("inorderTraversal", () => {
  test.each([
    {
      root: {
        val: 1,
        left: null,
        right: {
          val: 2,
          left: {
            val: 3,
            left: null,
            right: null,
          },
          right: null,
        },
      },
      want: [1, 3, 2],
    },
    {
      root: null,
      want: [],
    },
  ])("inorder traversal $root", ({ root, want }) => {
    expect(inorderTraversal(root)).toStrictEqual(want);
  });
});
