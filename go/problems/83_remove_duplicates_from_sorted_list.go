package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	temp := head
	for temp.Next != nil {
		if temp.Val == temp.Next.Val {
			temp.Next = temp.Next.Next
			continue
		}
		temp = temp.Next
	}

	return head
}
