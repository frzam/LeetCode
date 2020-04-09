/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    cycle := make(map[*ListNode]struct{})
    ok := false
    for head != nil{
        if _,ok = cycle[head];ok{
            return head
        }
         cycle[head] = struct{}{}
         head = head.Next
    }
    return nil
}