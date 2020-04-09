/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var merged *ListNode
	if l1.Val < l2.Val {
		merged = l1
		l1 = l1.Next
	} else {
		merged = l2
		l2 = l2.Next
	}
	head := merged
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			merged.Next = l1
			l1 = l1.Next
			merged = merged.Next
		} else {
			merged.Next = l2
			l2 = l2.Next
			merged = merged.Next

		}
	}
	for l1 != nil {
		merged.Next = l1
		l1 = l1.Next
		merged = merged.Next
	}
	for l2 != nil {
		merged.Next = l2
		l2 = l2.Next
		merged = merged.Next
	}
	return head
}