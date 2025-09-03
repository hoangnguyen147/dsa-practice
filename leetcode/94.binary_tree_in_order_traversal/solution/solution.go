package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	stack := []*TreeNode{root}
	res := make([]int, 0)
	for len(stack) > 0 {
		item := stack[len(stack)-1]
		if item.Left == nil {
			res = append(res, item.Val)
			stack = stack[:len(stack)-1]
			if item.Right != nil {
				stack = append(stack, item.Right)
			}
		} else {
			stack = append(stack, item.Left)
			item.Left = nil
		}
	}

	return res
}
