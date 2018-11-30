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

	h, w := NextInt(), NextInt()
	maze := make([][]int, h)
	//nodes := make([]*Node, 0)

	for i := 0; i < h; i++ {
		maze[i] = make([]int, w)
		line := NextInts()
		for j := 0; j < w; j++ {
			maze[i][j] = line[j]
		}
	}

	fmt.Println(CalcShortestPath(maze, w, h, 0, (h * w) - 1))
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
	q []int
	size int
}

func NewQueue(size int) *Queue {
	return &Queue{q: make([]int, 0, size)}
}

func (q *Queue) Enque(elem int) {
	q.q = append(q.q, elem)
	q.size++
}

func (q *Queue) EnqueRange(elems []int) {
	q.q = append(q.q, elems...)
	q.size += len(elems)
}

func (q *Queue) Deque() int {
	if q.size == 0 {
		return -1
	}

	res := q.q[0]
	q.q = q.q[1:]
	q.size--

	return res
}

func (q *Queue) Length() int {
	return q.size
}

func CalcShortestPath(maze [][]int, xSize int, ySize int, startPos int, endPos int) int {
	// 경로를 기억하도록 수정
	lengths := make([]int, xSize * ySize)
	//paths := make(map[int]int)
	history := make(map[int]bool)
	queue := NewQueue(0)
	queue.Enque(startPos)
	lengths[startPos] = 1
	for {
		pos := queue.Deque()
		if pos == endPos {
			break
		}

		nearPaths := FindNearPath(maze, xSize, ySize, pos, history)
		if len(nearPaths) == 0 {
			// 길이 없는 경우
			continue
		}
		for _, path := range nearPaths {
			// 이전 경로를 기록해둔다.
			//paths[path] = pos
			lengths[path] = lengths[pos] + 1
		}
		queue.EnqueRange(nearPaths)

		history[pos] = true
	}

	//return GetPathLength(paths, endPos, startPos)
	return lengths[endPos]
}

func GetPathLength(paths map[int]int, startTraverse int, endTraverse int) int {
	length := 0

	for {
		length++
		if startTraverse ==  endTraverse {
			break
		}
		startTraverse = paths[startTraverse]
	}

	return length
}

var directions = [4][2]int {
	{0, -1},
	{-1, 0}, {1, 0},
	{0, 1},
}

func FindNearPath(maze [][]int, xSize int, ySize int, targetPos int, history map[int]bool) []int {
	nearPoses := make([]int, 0)

	xPos := targetPos % xSize
	yPos := targetPos / ySize

	for _, direction := range directions {
		checkX := xPos + direction[0]
		checkY := yPos + direction[1]
		checkPos := (checkY * xSize) + checkX

		if (checkX >= 0) && (checkX < xSize) && (checkY >= 0) && (checkY < ySize) {
			if maze[checkY][checkX] == 1 && history[checkPos] == false {
				nearPoses = append(nearPoses, checkPos)
			}
		}
	}

	return nearPoses
}

