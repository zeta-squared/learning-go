package main

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	var head, node *ListNode
	sum, carry := 0, 0

	for l1 != nil && l2 != nil {
		sum = l1.Val + l2.Val + carry
		l1, l2 = l1.Next, l2.Next
		if sum >= 10 {
			sum %= 10
			carry = 1
		} else {
			carry = 0
		}

		if head == nil {
			head = &ListNode{
				Val: sum,
			}
			node = head
		} else {
			node.Next = &ListNode{
				Val: sum,
			}
			node = node.Next
		}
	}

	for l1 != nil {
		sum = l1.Val + carry
		l1 = l1.Next
		if sum >= 10 {
			sum %= 10
			carry = 1
		} else {
			carry = 0
		}

		if head == nil {
			head = &ListNode{
				Val: sum,
			}
			node = head
		} else {
			node.Next = &ListNode{
				Val: sum,
			}
			node = node.Next
		}
	}

	for l2 != nil {
		sum = l2.Val + carry
		l2 = l2.Next
		if sum >= 10 {
			sum %= 10
			carry = 1
		} else {
			carry = 0
		}

		if head == nil {
			head = &ListNode{
				Val: sum,
			}
			node = head
		} else {
			node.Next = &ListNode{
				Val: sum,
			}
			node = node.Next
		}
	}

	if carry > 0 {
		node.Next = &ListNode{
			Val: carry,
		}
	}

	return head
}
