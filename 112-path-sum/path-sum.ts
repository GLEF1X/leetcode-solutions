/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */

function hasPathSum(root: TreeNode | null, targetSum: number): boolean {
    if (targetSum === 0 && !root) return false

    const traverse = (node: TreeNode, currentSum: number = 0) => {
        currentSum += node.val

        if (!node.left && !node.right && currentSum === targetSum) {
            throw 5
        }

        if (node.left) {
            traverse(node.left, currentSum)
        }
        if (node.right) {
            traverse(node.right, currentSum)
        }
    }

    try {
        traverse(root)
    } catch (e: unknown) {
        if (typeof e === "number" && e === 5) {
            return true
        }
    }
    return false
};