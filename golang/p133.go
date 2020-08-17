package main

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

var nodeMap map[int]*Node
var nodeFlag map[int]bool

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	nodeMap = make(map[int]*Node)
	nodeFlag = make(map[int]bool)
	getGraph(node)
	return nodeMap[node.Val]
}

func getGraph(node *Node) {
	if nodeFlag[node.Val] {
		return
	}
	if _, ok := nodeMap[node.Val]; !ok {
		nodeMap[node.Val] = &Node{
			Val:       node.Val,
			Neighbors: []*Node{},
		}
	}
	nodeFlag[node.Val] = true
	for _, n := range node.Neighbors {
		if _, ok := nodeMap[n.Val]; !ok {
			nodeMap[n.Val] = &Node{
				Val:       n.Val,
				Neighbors: []*Node{},
			}
		}
		nodeMap[n.Val].Neighbors = append(nodeMap[n.Val].Neighbors, nodeMap[node.Val])
		getGraph(n)
	}
}