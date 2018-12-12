package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nodeNum int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d", &nodeNum)
	nodes := make([]*Node, nodeNum + 1)

	for n := 0; n < nodeNum; n++ {
		var n1, n2 int
		fmt.Fscanf(reader, "\n%d %d", &n1, &n2)
		if nodes[n1] == nil {
			nodes[n1] = NewNode(n1)
		}
		if nodes[n2] == nil {
			nodes[n2] = NewNode(n2)
		}
		nodes[n1].Connect(nodes[n2])
	}

	PrintTreeParents(nodes[1], nodeNum + 1)
}

type Node struct {
	id int
	neighbors []*Node
}

func (targetNode *Node) Connect(node *Node) {
	targetNode.neighbors = append(targetNode.neighbors, node)
	node.neighbors = append(node.neighbors, targetNode)
}

func NewNode(id int) *Node {
	return &Node{id: id, neighbors: make([]*Node, 0)}
}

func PrintTreeParents(root *Node, treeSize int) {
	cache := make([]int, treeSize)

	Traverse(root, nil, cache)

	for i := 2; i < treeSize; i++ {
		fmt.Println(cache[i])
	}
}

func Traverse(root *Node, prevNode *Node, cache []int) {
	if root == nil {
		return
	}

	if prevNode != nil {
		cache[root.id] = prevNode.id
	}

	for i := 0; i < len(root.neighbors); i++ {
		neighbor := root.neighbors[i]
		if neighbor == prevNode {
			continue
		}
		Traverse(neighbor, root, cache)
	}
}
