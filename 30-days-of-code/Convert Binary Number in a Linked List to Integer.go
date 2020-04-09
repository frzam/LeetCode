/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
    d := 0
    for head != nil{
         d = d * 2 + head.Val
         head = head.Next 
    }
    return d
}