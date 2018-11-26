package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var nodeNum, vertexNum, startNode int
	var nodes []*Node

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d %d", &nodeNum, &vertexNum, &startNode)
	nodes = make([]*Node, nodeNum + 1, nodeNum + 1)
	for i := 0; i < vertexNum; i++ {
		var startNodeId, endNodeId int
		fmt.Fscan(reader, &startNodeId)
		fmt.Fscan(reader, &endNodeId)

		if nodes[startNodeId] == nil {
			nodes[startNodeId] = NewNode(startNodeId)
		}
		if nodes[endNodeId] == nil {
			nodes[endNodeId] = NewNode(endNodeId)
		}

		nodes[startNodeId].neighbors = append(nodes[startNodeId].neighbors, nodes[endNodeId])
		nodes[endNodeId].neighbors = append(nodes[endNodeId].neighbors, nodes[startNodeId])
	}

	for _, node := range nodes {
		if node == nil {
			continue
		}

		// 작은 순대로 이웃을 정렬한다.
		sort.Sort(Nodes(node.neighbors))
	}

	DFS(nodeNum, startNode, nodes)
	fmt.Print("\n")
	CleanUp(nodes)
	BFS(nodeNum, startNode, nodes)
}

type Node struct {
	id int
	neighbors []*Node
	isVisited bool
}
type Nodes []*Node

func (nodes Nodes) Len() int {
	return len(nodes)
}

func (nodes Nodes) Less(i, j int) bool {
	return nodes[i].id <= nodes[j].id
}

func (nodes Nodes) Swap(i, j int) {
	nodes[i], nodes[j] = nodes[j], nodes[i]
}

func (node *Node) Visit() bool {
	if node.isVisited {
		return false
	}

	fmt.Printf("%d ", node.id)
	node.isVisited = true
	return true
}

func NewNode(id int) *Node {
	node := Node{id: id, neighbors: make([]*Node, 0, 5), isVisited: false}
	return &node
}

func CleanUp(nodes []*Node) {
	for _, node := range nodes {
		if node != nil {
			node.isVisited = false
		}
	}
}

func BFS(nodeNum int, startNode int,  nodes []*Node) {
	visitQueue := make([]*Node, 0, nodeNum * 2)
	visitQueue = append(visitQueue, nodes[startNode])
	for {
		if len(visitQueue) == 0 {
			return
		}

		target := visitQueue[0]
		visitQueue = visitQueue[1:]

		if isVisited := target.Visit(); isVisited {
			visitQueue = append(visitQueue, target.neighbors...)
		}
	}
}

func DFS(nodeNum int, startNode int, nodes []*Node) {
	visitStack := make([]*Node, 0)
	visitStack = append([]*Node{nodes[startNode]}, visitStack...)

	for {
		if len(visitStack) == 0 {
			return
		}

		target := visitStack[0]
		visitStack = visitStack[1:]

		if isVisited := target.Visit(); isVisited {
			neighbors := make([]*Node, len(target.neighbors))
			copy(neighbors, target.neighbors)
			visitStack = append(neighbors, visitStack...)
		}
	}
}

