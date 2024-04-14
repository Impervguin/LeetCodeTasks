// Time: 73.5%
// Memory: 11%

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    s := 0 // Sum of left leaves
    // Left node is a leaf, adding it's value
    if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
        s += root.Left.Val
    // Else there might be left leaves, so recursively check it
    } else if root.Left != nil {
        s += sumOfLeftLeaves(root.Left)
    }
    // There might be left leaves, so recursively check it
    if root.Right != nil {
        s += sumOfLeftLeaves(root.Right)
    }
    return s
}
