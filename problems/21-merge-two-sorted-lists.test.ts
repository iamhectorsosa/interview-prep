import { describe, expect, test } from "vitest";
import { mergeTwoLists, type ListNode } from "./21-merge-two-sorted-lists";

function toList(nums: Array<number>): ListNode | null {
  if (!nums.length) return null;
  return nums.reduceRight<ListNode | null>(
    (next, val) => ({ val, next }),
    null,
  );
}

describe("mergeTwoLists", () => {
  test.each([
    {
      list1: [1, 2, 4],
      list2: [1, 3, 4],
      want: [1, 1, 2, 3, 4, 4],
    },
    {
      list1: [],
      list2: [],
      want: [],
    },
    {
      list1: [],
      list2: [0],
      want: [0],
    },
  ])(
    "provided $list1 and $list2, we expect $want",
    ({ list1, list2, want }) => {
      expect(mergeTwoLists(toList(list1), toList(list2))).toStrictEqual(
        toList(want),
      );
    },
  );
});
