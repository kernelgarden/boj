package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var caseNum int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &caseNum)

	for c := 0; c < caseNum; c++ {
		var v, e int
		fmt.Fscanf(reader, "%d %d\n", &v, &e)
		nodes := make([]*Node, v + 1)
		// 편의를 위해서 0을 nil 로 채워 넣는다.
		nodes[0] = nil
		for i := 1; i <= v; i++ {
			nodes[i] = MakeNode(i)
		}

		for i := 0; i < e; i++ {
			var n1, n2 int
			fmt.Fscanf(reader, "%d %d\n", &n1, &n2)
			nodes[n1].Connect(nodes[n2])
		}

		if IsBipartiteGraph(nodes) {
			fmt.Println("YES")
		} else {
		 	fmt.Println("NO")
		}
	}
}

const (
	none  = 1
	red   = 2
	black = 3
)

type color int

 type Node struct {
 	id int
 	neighbor []*Node
 	color color
 }

func (node *Node) Connect(targetNode *Node) {
	if node == targetNode {
		return
	}
	node.neighbor = append(node.neighbor, targetNode)
	targetNode.neighbor = append(targetNode.neighbor, node)
}

func MakeNode(id int) *Node {
	return &Node{id: id, neighbor: make([]*Node, 0), color: none}
}

func chooseColor(prevColor color) color {
	if prevColor == red {
		return black
	} else {
		return red
	}
}

/*
func dspUtil(node *Node, toFillColor color, isVisited []bool) bool {
	// 이미 탐색했던 노드라면 색상의 모순 검출
	if isVisited[node.id] {
		// 이경우가 성립한다면 이분 그래프가 아니게 된다.
		if node.color != toFillColor {
			return false
		} else {
			return true
		}
	} else {
		isVisited[node.id] = true
	}

	isValid := true
	node.color = toFillColor
	for _, _node := range node.neighbor {
		isValid = dspUtil(_node, chooseColor(toFillColor), isVisited)
		if !isValid {
			break
		}
	}

	return isValid
}
*/
func dspUtil(node *Node, toFillColor color, isVisited []bool) bool {
	isValid := true
	node.color = toFillColor
	isVisited[node.id] = true
	for _, _node := range node.neighbor {
		// 인접 접점의 색이 칠해져 있는 경우 체크
		if _nodeColor := _node.color; _nodeColor != none {
			// 인접한 정점의 색이 같으면 이분 그래프가 아니다.
			if toFillColor == _nodeColor {
				isValid = false
				break
			}
		} else {
			// 색을 칠해나간다.
			if isVisited[_node.id] {
				continue
			}

			isValid = dspUtil(_node, chooseColor(toFillColor), isVisited)
		}
	}

	return isValid
}

func IsBipartiteGraph(nodes []*Node) bool {
	if len(nodes) == 1 {
		return false
	}
	isVisited := make([]bool, len(nodes) + 1)
	isBipartiteGraph := true
	for i := 1; i < len(nodes); i++ {
		if !isVisited[i] {
			isBipartiteGraph = dspUtil(nodes[1], color(black), isVisited)
		}
	}
	return isBipartiteGraph
}
