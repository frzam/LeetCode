/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
        temp := head
        for head.Next != nil && head.Next.Next != nil{
          
            head = head.Next.Next
            temp = temp.Next
        }
        if head.Next != nil{
            temp = temp.Next
        }

    return temp
}
