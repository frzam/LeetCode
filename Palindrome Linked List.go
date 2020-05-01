/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    cur := head
    fast :=head
    for fast != nil && fast.Next != nil{
        cur = cur.Next
        fast = fast.Next.Next
    }
    if fast != nil{
        cur = cur.Next
    }
    cur = reverse(cur)
    
    for cur != nil{
        if cur.Val != head.Val{
            return false
        }
        cur = cur.Next
        head = head.Next
    }
    return true
}

func reverse(head *ListNode)*ListNode{
    var prev *ListNode
    var next *ListNode
    for head != nil{
        next = head.Next
        head.Next = prev
        prev = head
        head = next
    }
    return prev
}
