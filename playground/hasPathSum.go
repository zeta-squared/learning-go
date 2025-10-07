package main

type SummedNode struct {
	node *TreeNode
	sum  int
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	hasPath := false
	if root == nil {
		return hasPath
	}
	stack := make([]SummedNode, 1)
	stack[0] = SummedNode{
		node: root,
		sum:  0,
	}

	for {
		if len(stack) == 0 {
			break
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node, sum := top.node, top.sum
		if node.Left == nil && node.Right == nil && sum+node.Val == targetSum {
			hasPath = true
			break
		}

		sum += node.Val
		if node.Right != nil {
			stack = append(stack, SummedNode{node: node.Right, sum: sum})
		}

		if node.Left != nil {
			stack = append(stack, SummedNode{node: node.Left, sum: sum})
		}
	}

	return hasPath
}
