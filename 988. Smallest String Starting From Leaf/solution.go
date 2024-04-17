// Time: 80.95%
// Memory: 7.14%

func findSmallestPath(root *TreeNode, nowPath []int) []int {
    if (root == nil) {
        return nil
    }
    if (root.Right == nil && root.Left == nil) {
        return append(nowPath, root.Val)
    }
    path1 := make([]int, len(nowPath), len(nowPath))
    path2 := make([]int, len(nowPath), len(nowPath))
    copy(path1, nowPath)
    copy(path2, nowPath)
    path1 = findSmallestPath(root.Left, append(path1, root.Val))
    path2 = findSmallestPath(root.Right, append(path2, root.Val))
    if (path1 == nil) {
        return path2
    } else if path2 == nil {
        return path1
    }
    for i, j := len(path1) - 1, len(path2) - 1; i >= 0 && j >= 0; {
        if path1[i] > path2[j] {
            return path2
        } else if path1[i] < path2[j] {
            return path1
        }
        i--
        j--
    }
    if (len(path1) > len(path2)) {
        return path2
    }
    return path1
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) string {
    res := findSmallestPath(root, []int{})
    strres := ""
    d := map[int]byte{}
    for i := 0; i < 26; i++ {
        d[i] = 'a' + byte(i)
    }
    for i := len(res) - 1; i >= 0; i-- {
        strres += string(d[res[i]])
    }
    return strres
}