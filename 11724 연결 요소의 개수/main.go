package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nodeNum, vertexNum int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &nodeNum)
	fmt.Fscan(reader, &vertexNum)

	nodes := make([]*Node, nodeNum, nodeNum)
	for i := 0; i < nodeNum; i++ {
		nodes[i] = NewNode(i)
	}

	for i := 0; i < vertexNum; i++ {
		var n1, n2 int
		fmt.Fscan(reader, &n1)
		fmt.Fscan(reader, &n2)

		nodes[n1 - 1].Connect(nodes[n2 - 1])
	}

	connectedNum := 0
	for {
		if len(nodes) == 0 {
			break
		}

		connectedIds := nodes[0].GetConnectedNodeId()

		for _, id := range connectedIds {
			nodes = Remove(nodes, id)
		}
		connectedNum++
	}

	fmt.Print(connectedNum)
}

func Remove(slice []*Node, target int) []*Node {
	targetIdx := -1
	for idx, node := range slice {
		if node == nil {
			continue
		}

		if node.id == target {
			targetIdx = idx
			break
		}
	}

	if targetIdx != -1 {
		slice[targetIdx] = slice[len(slice) - 1]
		slice = slice[:len(slice) - 1]
	}

	return slice
}

type Node struct {
	id int
	neighbors []*Node
	isVisited bool
}

func NewNode(id int) *Node {
	newNode := Node{id: id, neighbors: make([]*Node, 0)}
	return &newNode
}

func (startNode *Node) Connect(endNode * Node) {
	startNode.neighbors = append(startNode.neighbors, endNode)
	endNode.neighbors = append(endNode.neighbors, startNode)
}

func (node *Node) Visit() bool {
	if node.isVisited {
		return false
	}

	node.isVisited = true
	return true
}

func (startNode *Node) GetConnectedNodeId() []int {
	visitQueue := make([]*Node, 0)
	visitQueue = append(visitQueue, startNode)

	connectedNodeIds := make([]int, 0, 10)

	for {
		if len(visitQueue) == 0 {
			break
		}

		target := visitQueue[0]
		visitQueue = visitQueue[1:]

		if isNewlyVisited := target.Visit(); isNewlyVisited {
			connectedNodeIds = append(connectedNodeIds, target.id)
			visitQueue = append(visitQueue, target.neighbors...)
		}
	}

	return connectedNodeIds
}
