package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
*/
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	cloned := make([]*Node, 101)
	n := clone(node, cloned)
	return n
}

func clone(node *Node, cloned []*Node) *Node {
	if cloned[node.Val] != nil {
		return cloned[node.Val]
	}
	n := &Node{Val: node.Val, Neighbors: make([]*Node, 0, len(node.Neighbors))}
	cloned[node.Val] = n
	for i := range node.Neighbors {
		tmp := clone(node.Neighbors[i], cloned)
		n.Neighbors = append(n.Neighbors, tmp)
	}
	return n	
}
