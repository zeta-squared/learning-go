package main

func flatten(root *TreeNode) {
	stack := make([]*TreeNode, 1, 5)
	stack[0] = root
	var head *TreeNode

	for {
		if len(stack) == 0 {
			break
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node == nil {
			continue
		}

		stack = append(stack, node.Right, node.Left)
		if head.Right == nil {
			head.Right = node
			head.Right = head.Left
			head.Left = nil
		} else {
			head.Right = node
			head.Left = nil
			head = head.Right
		}
	}
}
