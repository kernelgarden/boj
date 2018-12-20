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

	fmt.Fscanf(reader, "\n")
	inOrders := make([]int, nodeNum)
	for i := 0; i < nodeNum; i++ {
		fmt.Fscanf(reader, "%d", &inOrders[i])
	}

	fmt.Fscanf(reader, "\n")
	postOrders := make([]int, nodeNum)
	for i := 0; i < nodeNum; i++ {
		fmt.Fscanf(reader, "%d", &postOrders[i])
	}

	tree := MakeTree(inOrders, postOrders)
	Traverse(tree)
}

func MakeTree(inOrders, postOrders []int) *Node {
	cache := PreProcess(inOrders)
	nodeNum := len(postOrders)
	root := MakeSubTree(inOrders, postOrders, nodeNum, cache, 0)
	return root
}

func MakeSubTree(subInOrders []int, subPostOrders []int, subLen int, cache map[int]int, factor int) *Node {
	if subLen == 1 {
		return MakeNode(subInOrders[0])
	}

	// factor는 cache를 스케일링해서 사용하기 위함이다.
	rootVal := subPostOrders[subLen - 1]
	rootPos := cache[rootVal] - factor

	root := MakeNode(rootVal)

	// 루트를 제외해야함
	//fmt.Printf("[Debug] rootVal: %v, len(subInOrders): %v\n", rootVal, len(subInOrders))
	left := subInOrders[:rootPos]
	right := subInOrders[rootPos + 1:]

	leftLen := len(left)
	rightLen := len(right)

	// 루트를 제외해야함
	//fmt.Printf("[Debug] leftLen: %v\n", leftLen)
	leftPost := subPostOrders[: leftLen]
	rightPost := subPostOrders[leftLen : subLen - 1]

	if leftLen > 0 {
		root.SetLeft(MakeSubTree(left, leftPost, leftLen, cache, factor))
	}

	if rightLen > 0 {
		root.SetRight(MakeSubTree(right, rightPost, rightLen, cache, factor + leftLen + 1))
	}

	return root
}

func PreProcess(inOrders []int) map[int]int {
	cache := make(map[int]int)
	for idx, v := range inOrders {
		cache[v] = idx
	}

	return cache
}

func Traverse(root *Node) {
	if root == nil {
		return
	}

	fmt.Printf("%d ", root.val)
	Traverse(root.left)
	Traverse(root.right)
}

type Bound struct {
	left int
	right int
}

func NewBound(left, right int) Bound {
	return Bound{left, right}
}

type Node struct {
	val int
	left, right *Node
	parent *Node
}

func MakeNode(val int) *Node {
	return &Node{val: val}
}

func (parent *Node) SetLeft(child *Node) {
	parent.left = child
	child.parent = parent
}

func (parent *Node) SetRight(child *Node) {
	parent.right = child
	child.parent = parent
}

func (n1 *Node) SetVal(val int) {
	n1.val = val
}

