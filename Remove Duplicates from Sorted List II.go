func deleteDuplicates(head *ListNode) *ListNode {
    cur := head
    dummy := &ListNode{
        Next:head,
    }
    prev := dummy
    for cur != nil{
        for cur.Next != nil && cur.Next.Val == cur.Val{
            cur = cur.Next
        }
        if prev.Next == cur{
            prev= prev.Next
        }else{
            prev.Next = cur.Next
        }
        cur = cur.Next
    }
    return dummy.Next
}
