/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(root *ListNode, k int) []*ListNode {
    // length of list
    head := root
    l := 0
    res := make([]*ListNode,k)
    for root !=nil{
        root = root.Next
        l ++
    }
    // No of items in the basket
    n := l /k
    extra := l %k
    for i:= 0;i<k;i++{
        top := new(ListNode)
        topHead := top
        s := n
        for head != nil && s > 0{
            top.Next = head
            head = head.Next
            top = top.Next
            s--
        }
        if extra > 0 && head != nil{
            top.Next = head
            head = head.Next
            top = top.Next
            extra --
        }
        top.Next = nil
        res[i] = topHead.Next
    }
    return res
}
