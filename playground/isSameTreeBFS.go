package main

func isSameTreeBFS(p, q *TreeNode) bool {
	pStack, qStack := make([]*TreeNode, 1, 5), make([]*TreeNode, 1, 5)
	pStack[0], qStack[0] = p, q
	isSame := true

	for {
		if len(pStack) != len(qStack) {
			isSame = false
			break
		}

		if len(pStack) == 0 {
			break
		}

		pNode, qNode := pStack[0], qStack[0]
		pStack, qStack = pStack[1:], qStack[1:]
		if pNode == nil && qNode == nil {
			continue
		}

		if (pNode == nil && qNode != nil) || (pNode != nil && qNode == nil) {
			isSame = false
			break
		}

		if pNode.Val != qNode.Val {
			isSame = false
			break
		}

		pStack = append(pStack, pNode.Right, pNode.Left)
		qStack = append(qStack, qNode.Right, qNode.Left)
	}

	return isSame
}
