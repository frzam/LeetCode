/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	merge := make(map[*ListNode]struct{})
	for headA != nil {
		merge[headA] = struct{}{}
		headA = headA.Next
	}
	ok := false
	for headB != nil {
		if _, ok = merge[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}