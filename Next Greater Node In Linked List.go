/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
    var res []int
    count := 0
    var temp *ListNode
    for head.Next != nil{
        count ++
        temp = head.Next
        for temp != nil && temp.Val<=head.Val{
            temp = temp.Next
        }
        if temp == nil{
            res = append(res, 0)
        }else{
            res = append(res, temp.Val)
        }
        head = head.Next
    }
    if count+1 >len(res){
        res = append(res, 0)
        count --
    }
    return res
}
