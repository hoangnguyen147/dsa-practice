package solution

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node struct {
	Val int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    if len(node.Neighbors) == 0 {
        return &Node{
            Val: node.Val,
            Neighbors: []*Node{},
        }
    }
    nodeMap := make(map[int]*Node)
    nodeMap[node.Val] = &Node{
        Val: node.Val,
        Neighbors: []*Node{},
    }
    stack := []*Node{node}
    for len(stack) > 0 {
        item := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        for _, nb := range item.Neighbors {
            if _, ok := nodeMap[nb.Val]; !ok {
                nodeMap[nb.Val] = &Node{
                    Val: nb.Val,
                    Neighbors: []*Node{},
                }
                stack = append(stack, nb)
            }
            nodeMap[item.Val].Neighbors = append(nodeMap[item.Val].Neighbors, nodeMap[nb.Val])
        }
    }
    return nodeMap[node.Val]
}
