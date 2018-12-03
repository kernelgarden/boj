package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var caseNum int
	fmt.Fscanf(reader, "%d\n", &caseNum)
	for i := 0; i < caseNum; i++ {
		var cycleSize int
		fmt.Fscanf(reader, "%d\n", &cycleSize)
		cycles := make([]*Node, 0)
		// 코딩상의 편의를 위해 0번 인덱스를 비워놓는다.
		cycles = append(cycles, nil)
		for cycle := 0; cycle < cycleSize; cycle++ {
			var linkTo int
			fmt.Fscanf(reader, "%d", &linkTo)
			newNode := MakeNode(cycle + 1, linkTo)
			cycles = append(cycles, newNode)
		}
		fmt.Fscanf(reader, "\n")
		fmt.Println(CalcCycleNum(cycles))
	}
}

type Node struct {
	id int
	linkTo int
	isVisited bool
}

func (node *Node) Visit() bool {
	if node.isVisited {
		return false
	}
	node.isVisited = true

	return true
}

func MakeNode(id int, linkTo int) *Node {
	return &Node{id: id, linkTo: linkTo, isVisited: false}
}

func Traverse(node *Node, cycles []*Node) []*Node {
	nodes := make([]*Node, 0)
	if node == nil || node.isVisited {
		return nil
	}

	startNodeId := node.id

	for {
		node.Visit()
		node = cycles[node.linkTo]

		if node == nil || node.id == startNodeId {
			break
		}
	}

	return nodes
}

func CalcCycleNum(cycles []*Node) int {
	cycleNum := 0

	for i := 1; i < len(cycles); i++ {
		if Traverse(cycles[i], cycles) != nil {
			cycleNum++
		}
	}

	return cycleNum
}
