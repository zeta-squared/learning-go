package main

func invertTreeBFS(root *TreeNode) *TreeNode {
	stack := make([]*TreeNode, 5)
	stack[0] = root

	for {
		if len(stack) == 0 {
			break
		}
		node := stack[0]
		stack = stack[1:]
		if node == nil {
			continue
		}

		// Push children of current node to stack even if they are nil
		stack = append(stack, node.Right, node.Left)

		// Switch children of current node
		tempLeft := node.Left
		node.Left = node.Right
		node.Right = tempLeft
	}

	return root
}
