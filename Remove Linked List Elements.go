/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    temp := head
    if head == nil{
        return head
    }
    for head.Next!= nil {
        if head.Next.Val == val && head.Next.Next != nil{
            head.Next = head.Next.Next
        }else if head.Next.Val == val && head.Next.Next == nil{
            head.Next = nil
        }else{
            head = head.Next
        }
    }
    if temp.Val == val{
        temp = temp.Next
    }
    return temp
}
