package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    stack := []*TreeNode{root}
    res := make([]int, 0)
    for len(stack) > 0 {
        item := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res = append(res, item.Val)
        if item.Right != nil {
            stack = append(stack, item.Right)
        }
        if item.Left != nil {
            stack = append(stack, item.Left)
        }
    }
    
    return res
}
