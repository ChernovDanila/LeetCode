/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int { 

    return counterDepth(root, 0)

}

func counterDepth(root *TreeNode, depth int) int{

    if root == nil {
        return depth
    }
    depth ++
    
    resLeft := counterDepth(root.Left, depth)
    resRight := counterDepth(root.Right, depth)
    
    return max(resLeft, resRight) 
    
}