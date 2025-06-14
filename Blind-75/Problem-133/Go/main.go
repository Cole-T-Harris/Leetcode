package main


// Definition for a Node.
type Node struct {
	Val int
	Neighbors []*Node
}


func cloneGraph(node *Node) *Node {
    visitedNodes := make(map[*Node]*Node)
	if node == nil {
		return nil
	}
	return cloneNode(node, &visitedNodes)
}

func cloneNode(node *Node, visitedNodes *map[*Node]*Node) *Node {
		if node == nil {
		return nil
	}
	if  cloned, found:= (*visitedNodes)[node]; found {
		return cloned
	}
	newNode := Node{
		Val: node.Val,
		Neighbors: []*Node{},
	}
	(*visitedNodes)[node] = &newNode
		for _, neighborNode := range node.Neighbors {
		newNeighborNode := cloneNode(neighborNode, visitedNodes)
		newNode.Neighbors = append(newNode.Neighbors, newNeighborNode)
	}
	return &newNode
}