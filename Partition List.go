/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    less :=new(ListNode)
    more :=new(ListNode)
    
    lessHead := less
    moreHead := more
    
    for head != nil{
        if head.Val <x{
            less.Next = head
            less = less.Next
        }else{
            more.Next = head
            more = more.Next
        }
        head = head.Next
    }
    less.Next =moreHead.Next
    more.Next = nil
    return lessHead.Next
iiii}
