/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

function mergeTwoLists(list1: ListNode | null, list2: ListNode | null): ListNode | null {
    const [list1Items, list2Items] = [[], []]
    
    if (list1) {
        let current = list1
        while (current) {
            list1Items.push(current.val)
            current = current.next
        }
    }

    if (list2) {
        let current = list2
        while (current) {
            list1Items.push(current.val)
            current = current.next
        }
    }
    
    const result = [...list1Items, ...list2Items]
    result.sort((a, b) => a - b)

    const root = new ListNode()

    if (result.length) {
        const firstItem = result.shift()
        root.val = firstItem
    } else {
        return null
    }

    let current = root
    while (result.length) {
        const item = result.shift()
        current.next = new ListNode(item)
        current = current.next
    }

    return root
    
};