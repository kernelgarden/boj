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
	fmt.Fscanf(reader, "%d", &nodeNum)

	nodes := make([]*Node, nodeNum + 1)
	nodes[1] = NewNode(1)

	for i := 0; i < nodeNum; i++ {
		var p, c, w int
		fmt.Fscanf(reader, "\n%d %d %d", &p, &c, &w)

		if nodes[p] == nil {
			nodes[p] = NewNode(p)
		}
		if nodes[c] == nil {
			nodes[c] = NewNode(c)
		}
		nodes[p].AddChild(nodes[c], w)
	}

	fmt.Println(GetLongestLength(nodes[1]))
}

func GetLongestLength(root *Node) int {
	longestPath := 0
	pathToRoot := Traverse(root, &longestPath)
	if pathToRoot > longestPath  {
		return pathToRoot
	} else {
		return longestPath
	}
}

func Traverse(root *Node, currentLongestPath *int) int {
	if root == nil {
		return 0
	}

	paths := make([]int, 0)

	for i := 0; i < len(root.children); i++ {
		length := Traverse(root.children[i], currentLongestPath) + root.weights[i]
		paths = append(paths, length)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(paths)))

	first := 0
	second := 0
	if len(paths) >= 2 {
		first = paths[0]
		second = paths[1]
	} else if len(paths) == 1 {
		first = paths[0]
	}

	totalPath := first + second
	if totalPath > *currentLongestPath  {
		*currentLongestPath = totalPath
	}

	//fmt.Printf("[Debug] => root: %d, path: %d, first: %d, second: %d\n", root.val, totalPath, first, second)

	return first
}

type Node struct {
	val int
	children []*Node
	weights []int
}

func NewNode(val int) *Node {
	return &Node{val: val, children: make([]*Node, 0), weights: make([]int, 0)}
}

func (parent *Node)AddChild(child *Node, weight int) {
	parent.children = append(parent.children, child)
	parent.weights = append(parent.weights, weight)
}

/* BinTree Version
type Node struct {
	val int
	left, right *Node
	lWeight, rWeight int
}

func NewNode(val int) *Node {
	return &Node{val: val}
}

func (parent *Node)AddChild(child *Node, weight int) {
	if parent.left == nil {
		parent.SetLeft(child, weight)
	} else {
		parent.SetRight(child, weight)
	}
}

func (parent *Node)SetLeft(child *Node, weight int) {
	parent.left = child
	parent.lWeight = weight
}

func (parent *Node)SetRight(child *Node, weight int) {
	parent.right = child
	parent.rWeight = weight
}
*/
