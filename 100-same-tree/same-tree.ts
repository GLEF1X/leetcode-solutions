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

function isSameTree(p: TreeNode | null, q: TreeNode | null): boolean {
    if (!p && !q) {
        return true
    }

    if (!!p !== !!q) {
        return false
    }

    const queueP = [p] as TreeNode[]
    const queueQ = [q] as TreeNode[]
    while(queueP.length && queueQ.length) {
        const [qNode, pNode] = [queueQ.shift(), queueP.shift()]
        if (!!qNode !== !!pNode) {
            return false
        }
        if (!qNode) return true

        if (qNode.val !== pNode.val) {
            return false
        }
        if (!!qNode.left !== !!pNode.left) return false
        if (!!qNode.right !== !!pNode.right) return false

        if (qNode.left) {
            queueQ.push(qNode.left)
            queueP.push(pNode.left)
        }

        if (qNode.right) {
            queueQ.push(qNode.right)
            queueP.push(pNode.right)
        }
    }

    return true
};