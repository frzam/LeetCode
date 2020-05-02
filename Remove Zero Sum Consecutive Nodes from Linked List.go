/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
    prefix := 0
    top := &ListNode{
        Next:head,
    }
    seen := make(map[int]*ListNode)
    seen[0] = top
    for cur := top;cur != nil;cur = cur.Next{
        prefix += cur.Val
        seen[prefix]=cur
    }
    prefix = 0
    for cur := top;cur != nil;cur = cur.Next{
        prefix += cur.Val
        if p,ok := seen[prefix];ok{
              cur.Next = p.Next
       }
    }
    return top.Next
}
