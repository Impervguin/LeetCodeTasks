// Time: 78.59%
// Memory: 36.67

func sumNumbersSum(root *TreeNode, rootSum int) int {
    s := 0 // sum of numbers
    rootSum = rootSum * 10 + root.Val // add current number to rootSum
    // if leaf, then add current sum to sum
    if root.Left == nil && root.Right == nil {
        s += rootSum
    }
    // recursively check left leaves
    if root.Left != nil {
        s += sumNumbersSum(root.Left, rootSum)
    }
    // recursively check right leaves
    if root.Right != nil {
        s += sumNumbersSum(root.Right, rootSum)
    }
    return s
}


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    return sumNumbersSum(root, 0)
}