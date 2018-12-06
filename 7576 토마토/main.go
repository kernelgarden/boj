package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	Blank = -1
	Unfinished = 0
	Finished = 1
)

func main() {
	var w, h int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d", &w, &h)

	tomatoes := make([][]int, h)
	startPoses := make([]int, 0)

	for i := 0; i < h; i++ {
		fmt.Fscanf(reader, "\n")
		tomatoes[i] = make([]int, w)
		for j := 0; j < w; j++ {
			fmt.Fscanf(reader, "%d", &tomatoes[i][j])
			if tomatoes[i][j] == Finished {
				startPoses = append(startPoses, getPos(j, i, w))
			}
		}
	}

	fmt.Println(getFinishedDay(tomatoes, w, h, startPoses))
}

func getFinishedDay(tomatoes [][]int, w int, h int, startPoses []int) int {
	day := 0
	queue := make([]int, 0)
	queue = append(queue, startPoses...)
	isVisited := make([]bool, w * h)
	for i := 0; i < w * h; i++ {
		isVisited[i] = false
	}
	leftTargets := len(startPoses)

	for {
		if len(queue) == 0 {
			break
		}

		first := queue[0]
		queue = queue[1:]

		// 이미 방문했던 노드는 체크하지 않아도 된다.
		if isVisited[first] {
			continue
		} else {
			isVisited[first] = true
		}

		affectedTomatoes := AffectTomato(tomatoes, w, h, first)
		for _, affectedTomato := range affectedTomatoes {
			queue = append(queue, affectedTomato)
		}

		if leftTargets == 1 {
			day++
			leftTargets = len(queue)
		} else {
			leftTargets--
		}
	}

	if checkRemainTomatoes(tomatoes, w, h) {
		return -1
	}

	return day - 1
}

func getPos(x, y, w int) int {
	return (y * w) + x
}

func getXY(pos, x int) (int, int) {
	return pos % x, pos / x
}

var directions = [4][2]int {
	{0, -1}, {-1, 0}, {1, 0}, {0, 1},
}

// 영향 받은 노드들의 위치를 리턴한다.
func AffectTomato(tomatoes [][]int, w int, h int, target int) ([]int) {
	affectedNodes := make([]int, 0)
	x, y := getXY(target, w)
	for _, direction := range directions {
		checkX := x + direction[0]
		checkY := y + direction[1]

		if (checkX >= 0) && (checkX < w) && (checkY >= 0) && (checkY < h) {
			if tomatoes[checkY][checkX] == 0 {
				tomatoes[checkY][checkX] = 1
				affectedNodes = append(affectedNodes, getPos(checkX, checkY, w))
			}
		}
	}

	return affectedNodes
}

func checkRemainTomatoes(tomatoes [][ ]int, w int, h int) bool {
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if tomatoes[i][j] == 0 {
				return true
			}
		}
	}

	return false
}