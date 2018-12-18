package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var nodeNum int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &nodeNum)

	nodes := make([]*Node, nodeNum + 1)
	var root *Node

	//for i := 0; i < nodeNum; i++ {
	for {
		var p int
		fmt.Fscanf(reader, "%d", &p)

		if p == 0 {
			break
		}

		if nodes[p] == nil {
			nodes[p] = MakeNode(p)
		}

		if root == nil {
			root = nodes[p]
		}

		for {
			var c, w int
			fmt.Fscanf(reader, "%d %d", &c, &w)
			if c == -1 {
				break
			}

			if nodes[c] == nil {
				nodes[c] = MakeNode(c)
			}

			nodes[p].Connect(nodes[c], w)
		}
	}

	fmt.Println(getLongestPath(root))
}

type Node struct {
	val int
	neighbors []*Node
	weights []int
	isVisited bool
}

func MakeNode(val int) *Node {
	return &Node{val: val, neighbors: make([]*Node, 0), weights: make([]int, 0), isVisited: false}
}

func (n1 *Node) Connect(n2 *Node, weight int) {
	n1.neighbors = append(n1.neighbors, n2)
	n1.weights = append(n1.weights, weight)
}

func getLongestPath(root *Node) int {
	longestPath := 0

	lengthToRoot := traverse(root, &longestPath)

	if lengthToRoot >= longestPath {
		return lengthToRoot
	} else {
		return longestPath
	}
}

func traverse(root *Node, longestPath *int) int {
	if root == nil {
		return 0
	}
	root.isVisited = true

	lengths := make([]int, 0)
	for i := 0; i < len(root.neighbors); i++ {
		if root.neighbors[i].isVisited {
			continue
		}
		lengths = append(lengths, traverse(root.neighbors[i], longestPath) + root.weights[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(lengths)))

	first := 0
	second := 0

	if len(lengths) >= 2 {
		first = lengths[0]
		second = lengths[1]
	} else if len(lengths) == 1 {
		first = lengths[0]
	}

	totalPath := first + second
	if totalPath > *longestPath {
		*longestPath = totalPath
	}

	//fmt.Printf("[Debug] root: %v, path: %v, first: %v, second: %v\n", root.val, totalPath, first, second)

	return first
}
