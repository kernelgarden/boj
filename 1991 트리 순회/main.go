package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nodeNum int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &nodeNum)

	tmpMap := make(map[rune]*Node)

	treeRoot := rune(' ')

	for i := 0; i < nodeNum; i++ {
		var root, left, right rune
		fmt.Fscanf(reader, "%c %c %c\n", &root, &left, &right)

		if _, has := tmpMap[root]; !has {
			tmpMap[root] = NewNode(root)
		}

		if treeRoot == rune(' ') {
			treeRoot = root
		}

		if left != '.' {
			if _, has := tmpMap[left]; !has {
				tmpMap[left] = NewNode(left)
			}

			tmpMap[root].AddLeft(tmpMap[left])
		}
		if right != '.' {
			if _, has := tmpMap[right]; !has {
				tmpMap[right] = NewNode(right)
			}

			tmpMap[root].AddRight(tmpMap[right])
		}
	}

	TraverseBefore(tmpMap[treeRoot])
	fmt.Println()
	TraverseMid(tmpMap[treeRoot])
	fmt.Println()
	TraverseAfter(tmpMap[treeRoot])
	fmt.Println()
}

type Node struct {
	id rune
	left *Node
	right *Node
}

func NewNode(id rune) *Node{
	return &Node{id: id, left: nil, right: nil}
}

func (target *Node) AddLeft(n1 *Node) {
	target.left = n1
}

func (target *Node) AddRight(n1 *Node) {
	target.right = n1
}

func TraverseBefore(startNode *Node) {
	if startNode == nil {
		return
	}

	fmt.Printf("%c", startNode.id)
	TraverseBefore(startNode.left)
	TraverseBefore(startNode.right)
}

func TraverseMid(startNode *Node) {
	if startNode == nil {
		return
	}

	TraverseMid(startNode.left)
	fmt.Printf("%c", startNode.id)
	TraverseMid(startNode.right)
}

func TraverseAfter(startNode *Node) {
	if startNode == nil {
		return
	}

	TraverseAfter(startNode.left)
	TraverseAfter(startNode.right)
	fmt.Printf("%c", startNode.id)
}
