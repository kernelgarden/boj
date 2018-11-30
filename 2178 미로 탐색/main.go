package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func NextInt() int {
	scanner.Scan()
	r := 0
	for _, c := range scanner.Bytes() {
		r *= 10
		r += int(c - '0')
	}
	return r
}

func NextInts() []int {
	scanner.Scan()
	res := make([]int, 0)
	for _, c := range scanner.Bytes() {
		res = append(res, int(c - '0'))
	}
	return res
}

func main() {
	scanner.Split(bufio.ScanWords)
	reader := bufio.NewReader(os.Stdin)

	//h, w := NextInt(), NextInt()
	var h, w int
	fmt.Fscanf(reader, "%d", &h)
	fmt.Fscanf(reader, "%d\n", &w)
	maze := make([][]int, h)
	nodes := make([]*Node, h * w)

	for i := 0; i < h; i++ {
		maze[i] = make([]int, w)
		//line := NextInts()
		for j := 0; j < w; j++ {
			//maze[i][j] = line[j]
			n, _ := reader.ReadByte()
			maze[i][j] = int(n - '0')
		}
		reader.ReadByte()
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			pos := (i * w) + j
			nearPoses := GetNearPoses(maze, w, h, j, i)
			nodes[pos] = NewNode(pos)
			nodes[pos].neighbors = append(nodes[pos].neighbors, nearPoses...)
		}
	}

	fmt.Println(CalcShortestPath(nodes, 0, (h * w) - 1))
}

type Node struct {
	pos int
	neighbors []int
	isVisited bool
}

func NewNode(pos int) *Node {
	return &Node{pos: pos, neighbors: make([]int, 0, 4), isVisited: false}
}

type Queue struct {
	q []*Node
	size int
}

func NewQueue(size int) *Queue {
	return &Queue{q: make([]*Node, 0, size)}
}

func (q *Queue) Enque(elem *Node) {
	q.q = append(q.q, elem)
	q.size++
}

func (q *Queue) EnqueRange(elems []*Node) {
	q.q = append(q.q, elems...)
	q.size += len(elems)
}

func (q *Queue) Deque() *Node {
	if q.size == 0 {
		return nil
	}

	res := q.q[0]
	q.q = q.q[1:]
	q.size--

	return res
}

func (q *Queue) Length() int {
	return q.size
}

func CalcShortestPath(nodes []*Node, startNode int, endNode int) int {
	// 경로를 기억하도록 수정
	lengths := make([]int, len(nodes))
	queue := NewQueue(0)
	queue.Enque(nodes[startNode])
	lengths[startNode] = 1
	for {
		pos := queue.Deque()
		if pos.isVisited {
			continue
		}

		if pos.pos == endNode {
			break
		}

		if len(pos.neighbors) == 0 {
			// 길이 없는 경우
			continue
		}

		pos.isVisited = true
		for _, neighbor := range pos.neighbors {
			if !nodes[neighbor].isVisited {
				queue.Enque(nodes[neighbor])
				lengths[nodes[neighbor].pos] = lengths[pos.pos] + 1
			}
		}
	}

	return lengths[nodes[endNode].pos]
}

var directions = [4][2]int {
	{0, -1},
	{-1, 0}, {1, 0},
	{0, 1},
}

func GetNearPoses(maze [][]int, xSize int, ySize int, xPos int, yPos int) []int {
	nearPoses := make([]int, 0)

	for _, direction := range directions {
		checkX := xPos + direction[0]
		checkY := yPos + direction[1]
		checkPos := (checkY * xSize) + checkX

		if (checkX >= 0) && (checkX < xSize) && (checkY >= 0) && (checkY < ySize) {
			if maze[checkY][checkX] == 1 {
				nearPoses = append(nearPoses, checkPos)
			}
		}
	}

	return nearPoses
}
