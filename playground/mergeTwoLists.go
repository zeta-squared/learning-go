package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	var node *ListNode
	for list1 != nil && list2 != nil {
		l1val, l2val := list1.Val, list2.Val
		switch {
		case l1val <= l2val:
			if head == nil {
				head = list1
				node = head
			} else {
				node.Next = list1
				node = node.Next
			}
			list1 = list1.Next
		default:
			if head == nil {
				head = list2
				node = head
			} else {
				node.Next = list2
				node = node.Next
			}
			list2 = list2.Next
		}
	}

	for list1 != nil {
		if head == nil {
			head = list1
			node = head
		} else {
			node.Next = list1
			node = node.Next
		}
		list1 = list1.Next
	}

	for list2 != nil {
		if head == nil {
			head = list2
			node = head
		} else {
			node.Next = list2
			node = node.Next
		}
		list2 = list2.Next
	}

	return head
}
