// Time: 100%
// Memory: 100%

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
    if (root == nil) {
        return root
    }
        
    if (depth == 1) {
        newRoot := TreeNode{}
        newRoot.Val = val
        newRoot.Left = root;
        return &newRoot;
    } else if (depth == 2) {
        newLeft, newRight := TreeNode{}, TreeNode{}
        newLeft.Val = val
        newRight.Val = val
        newLeft.Left = root.Left
        newRight.Right = root.Right
        root.Left = &newLeft
        root.Right = &newRight
        return root
    }
    addOneRow(root.Left, val, depth - 1)
    
    addOneRow(root.Right, val, depth - 1)
    return root;
}