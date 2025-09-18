export type ListNode = {
  val: number;
  next: ListNode | null;
};

export function deleteDuplicates(head: ListNode | null): ListNode | null {
  if (head === null || head.next === null) {
    return head;
  }

  let current = head;

  while (current.next !== null) {
    if (current.val === current.next.val) {
      current.next = current.next.next;
      continue;
    }
    current = current.next;
  }

  return head;
}
