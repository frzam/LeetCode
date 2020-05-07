/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
    var parentX *TreeNode
    var parentY *TreeNode
    var xD, yD int

func isCousins(root *TreeNode, x int, y int) bool {
    xD,yD = -1, -1
    getDepthAndParent(root, x ,y , 0, nil)
    return xD == yD && parentX != parentY
}

func getDepthAndParent(root *TreeNode, x int, y int, depth int,parent *TreeNode){
    if root == nil{
        return
    }
    if root.Val == x{
        parentX = parent
        xD = depth
    }else if root.Val == y{
        parentY = parent
        yD = depth
    }else{
        getDepthAndParent(root.Left, x , y , depth +1, root)
        getDepthAndParent(root.Right, x , y, depth +1, root)
    }
}
