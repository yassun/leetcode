package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	currentNode := head
	for currentNode != nil && currentNode.Next != nil {
		if currentNode.Val == currentNode.Next.Val {
			currentNode.Next = currentNode.Next.Next
		} else {
			currentNode = currentNode.Next
		}

	}
	return head
}
