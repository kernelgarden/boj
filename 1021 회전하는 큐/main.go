package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var queueSize, operationNum int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &queueSize)
	fmt.Fscan(reader, &operationNum)

	operations := make([]int, operationNum)
	for i := 0; i < operationNum; i++ {
		fmt.Fscan(reader, &operations[i])
	}

	result := Simulate(operations, queueSize)
	fmt.Println(result)
}

type RotateQueue struct {
	queue []int
	queueSize int
	front int
}

func (queue *RotateQueue) Left() {
	queue.front = (queue.front + 1) % queue.queueSize
}

func (queue *RotateQueue) Right() {
	if queue.front == 0 {
		queue.front = queue.queueSize - 1
	} else {
		queue.front = queue.front - 1
	}
}

func (queue *RotateQueue) Deque() int {
	ret := queue.queue[queue.front]

	// quque를 조정 한 후 queueSize를 줄이기 전에 front를 조정합니다!
	queue.queue = append(queue.queue[ : queue.front], queue.queue[queue.front + 1 : ]...)
	if queue.front == queue.queueSize - 1 {
		queue.front = 0
	}
	queue.queueSize = queue.queueSize - 1

	return ret
}

func NewRotateQueue(queueSize int) *RotateQueue{
	newRotateQueue := RotateQueue{queue: make([]int, queueSize, queueSize), queueSize: queueSize, front: 0}
	for i := 0; i < queueSize; i++ {
		newRotateQueue.queue[i] = i + 1
	}
	return &newRotateQueue
}

func RotateLeft(rotateQueue *RotateQueue, count int) {
	for i := 0; i < count ; i++ {
		rotateQueue.Left()
	}
}

func RotateRight(rotateQueue *RotateQueue, count int) {
	for i := 0; i < count; i++ {
		rotateQueue.Right()
	}
}

func Simulate(operations []int, queueSize int) int {
	history := make(map[string]int)
	history["left"] = 0
	history["right"] = 0

	rotateQueue := NewRotateQueue(queueSize)

	for {
		if len(operations) == 0 {
			return history["left"] + history["right"]
		}

		target := operations[0]
		operations = operations[1:]

		operation, dist := GetShortestDist(rotateQueue, target)

		if operation == "left" {
			RotateLeft(rotateQueue, dist)
		} else if operation == "right" {
			RotateRight(rotateQueue, dist)
		}
		history[operation] += dist
		rotateQueue.Deque()
	}
}

func FindTargetPos(rotateQueue *RotateQueue, target int) int {
	ret := -1

	for i := 0; i < rotateQueue.queueSize; i++ {
		if rotateQueue.queue[i] == target {
			ret = i
			break
		}
	}

	return ret
}

func GetShortestDist(rotateQueue *RotateQueue, target int) (string, int) {
	targetPos := FindTargetPos(rotateQueue, target)

	dist :=  int(math.Abs(float64(targetPos - rotateQueue.front)))
	reverseDist := rotateQueue.queueSize - dist

	if dist <= reverseDist {
		if targetPos >= rotateQueue.front {
			return "left", dist
		} else {
			return "right", dist
		}
	} else {
		if targetPos >= rotateQueue.front {
			return "right", reverseDist
		} else {
			return "left", reverseDist
		}
	}
}
