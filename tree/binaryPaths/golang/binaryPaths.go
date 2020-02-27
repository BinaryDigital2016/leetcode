/*
给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归 
func binaryTreePaths(root *TreeNode) []string {
    if root == nil{
        return []string{}
    }

    paths := make([]string,0)
    constructPath(root, "", &paths)
    return paths
}

// slice需传指针 否则append导致底层数组重新分配时会分配新的数组
func constructPath(node *TreeNode, path string, paths *[]string){
    path += fmt.Sprintf("%d", node.Val)
    if node.Left == nil && node.Right == nil{
        *paths = append(*paths, path)
        return
    }

    path += "->"
    if node.Left != nil {
        //path += fmt.Sprintf("%d", node.Left.Val)
        constructPath(node.Left, path, paths)
    }
    if node.Right != nil {
        //path += fmt.Sprintf("%d", node.Right.Val)
        constructPath(node.Right, path, paths)
    }
}

// // 迭代 
// func binaryTreePaths(root *TreeNode) []string {
//     if root == nil{
//         return []string{}
//     }

//     s := make([]pathInfo,0)
//     s = append(s,pathInfo{root,""})
//     paths := make([]string,0)
//     for len(s) > 0 {
//         info := s[0]
//         info.path += fmt.Sprintf("%d", info.node.Val)
//         s = s[1:]
//         if info.node.Left == nil && info.node.Right == nil {
//             paths = append(paths, info.path)
//             continue
//         }

//         info.path += "->"
//         if info.node.Left != nil {
//             s = append(s, pathInfo{info.node.Left, info.path})
//         }
//         if info.node.Right != nil {
//             s = append(s, pathInfo{info.node.Right, info.path})
//         }
//     }
    
//     return paths
// }

// type pathInfo struct{
//     node *TreeNode
//     path string
// }
