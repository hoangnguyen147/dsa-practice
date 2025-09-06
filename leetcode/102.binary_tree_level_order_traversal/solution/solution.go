package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    queue := []*TreeNode{root}
    res := [][]int{}
    
    for len(queue) > 0 {
        currLen := len(queue)
        currLevelVal := make([]int, 0, currLen)

        for i := 0; i < currLen; i++ {
            node := queue[0]
            queue = queue[1:]
            currLevelVal = append(currLevelVal, node.Val)
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }

            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        
        res = append(res, currLevelVal)
    }
    
    return res
}
