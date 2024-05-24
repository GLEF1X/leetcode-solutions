/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := countLeftDepthIterative(root)
	rightDepth := countRightDepthIterative(root)

	if leftDepth == rightDepth {
		return IntPow(2, leftDepth) - 1
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func countLeftDepth(node *TreeNode) int {
	if node.Left == nil {
		return 1
	}

	return 1 + countLeftDepth(node.Left)
}

func countLeftDepthIterative(node *TreeNode) int {
	var count int
	currentNode := node
	for currentNode != nil {
		count++
		currentNode = currentNode.Left
	}

	return count
}

func countRightDepth(node *TreeNode) int {
	if node.Right == nil {
		return 1
	}

	return 1 + countRightDepth(node.Right)
}

func countRightDepthIterative(node *TreeNode) int {
	var count int
	currentNode := node
	for currentNode != nil {
		count++
		currentNode = currentNode.Right
	}

	return count
}

func IntPow(base, exp int) int {
	result := 1
	for {
		if exp&1 == 1 {
			result *= base
		}
		exp >>= 1
		if exp == 0 {
			break
		}
		base *= base
	}

	return result
}
