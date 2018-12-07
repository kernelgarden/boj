package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, p int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d", &a, &p)
	fmt.Println(GetNoneCycleNum(a, p))
}

type Node struct {
	idx int
	val int
}

func NewNode(idx, val int) *Node {
	return &Node{idx: idx, val: val}
}

func GetNoneCycleNum(a, p int) int {
	cache := make(map[int]int)

	startNode := NewNode(0, a)
	curNode := startNode
	cache[curNode.val] = startNode.idx

	for {
		nextNode := calcNext(curNode.val, p, curNode.idx + 1)
		curNode = nextNode

		// 사이클을 발견했다.
		if _, has := cache[curNode.val]; has {
			break
		} else {
			cache[curNode.val] = curNode.idx
		}
	}

	return cache[curNode.val]
}

func calcNext(cur int, p int, nextIdx int) *Node {
	val := 0
	for {
		if cur == 0 {
			break
		}

		num := cur % 10
		cur /= 10
		val += pow(num, p)
	}

	return NewNode(nextIdx, val)
}

func pow(x, y int) int {
	val := x
	for i := 1; i < y; i++ {
		val *= x
	}
	return val
}


