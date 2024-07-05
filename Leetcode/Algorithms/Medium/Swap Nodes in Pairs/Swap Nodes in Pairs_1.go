/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	temp := head.Next.Next
	head.Next.Next = head
	head = head.Next
	head.Next.Next = swapPairs(temp)

	return head
}