import { describe, expect, test } from "vitest";
import {
  deleteDuplicates,
  type ListNode,
} from "./83-remove-duplicates-from-sorted-list";

function toList(nums: Array<number>): ListNode | null {
  if (!nums.length) return null;
  return nums.reduceRight<ListNode | null>(
    (next, val) => ({ val, next }),
    null,
  );
}

describe("deleteDuplicates", () => {
  test.each([
    {
      head: [1, 1, 2],
      want: [1, 2],
    },
    {
      head: [1, 1, 2, 3, 3],
      want: [1, 2, 3],
    },
    {
      head: [1, 1, 1],
      want: [1],
    },
  ])("provided $head, we expect $want", ({ head, want }) => {
    expect(deleteDuplicates(toList(head))).toStrictEqual(toList(want));
  });
});
