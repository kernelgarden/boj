package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var size int
	fmt.Fscanf(reader, "%d\n", &size)
	nodes := make([][]int, size)

	for i := 0; i <  size; i++ {
		nodes[i] = make([]int, size)
		for j := 0; j < size; j++ {
			fmt.Fscanf(reader, "%d", &nodes[i][j])
		}
		fmt.Fscanf(reader, "\n")
	}

	makeConnectMap(nodes, size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%v ", nodes[i][j])
		}
		fmt.Println()
	}

	/*
	shortestDist := math.MaxInt32
	findShortestDist(&shortestDist)
	fmt.Println(shortestDist)
	*/
}

func makeConnectMap(nodes [][]int, size int) {
	from := make([]int, size * size)
	to := make(map[int]map[int]bool)

	for i := 0; i < size * size; i++ {
		from[i] = i
		to[i] = map[int]bool{i: true}
	}

	for i := 0; i < size * size; i++ {
		connectedNodes, seaNodes := GetConnectedNodes(nodes, size, i)

		for _, connectedNode := range connectedNodes {
			a, b := from[connectedNode], from[i]

			if a == b {
				continue
			}

			if len(to[a]) < len(to[b]) {
				a, b = b, a
			}

			for v := range to[b] {
				from[v] = a
				to[a][v] = true
			}

			delete(to, b)
		}

		for _, seaNode := range seaNodes {
			from[seaNode] = 0
			delete(to, seaNode)
		}
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			pos := (i * size) + j
			if from[pos] != 0 {
				nodes[i][j] = from[pos]
			}
		}
	}
}

var directions = [4][2]int {
	{0, -1}, {-1, 0}, {1, 0}, {0, 1},
}

func GetConnectedNodes(nodes [][]int, size int, targetPos int) ([]int, []int) {
	connectedNodes := make([]int, 0)
	unConnectedNodes := make([]int, 0)

	targetX := targetPos % size
	targetY := targetPos / size

	for _, direction := range directions {
		checkX := direction[0] + targetX
		checkY := direction[1] + targetY

		if (checkX >= 0) && (checkX < size) && (checkY >= 0) && (checkY < size) {
			if nodes[checkY][checkX] == 1 {
				connectedNodes = append(connectedNodes, (checkY * size) + checkX)
			} else {
				unConnectedNodes = append(unConnectedNodes, (checkY * size) + checkX)
			}
		}
	}

	return connectedNodes, unConnectedNodes
}

func findShortestDist(nodes [][]int, size int, shortestDist *int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			// 탐색 시작 노드가 바다면 패스
			if nodes[i][j] == 0 {
				continue
			}

			startGroupId := nodes[i][j]

			queue := make([]int, 0)
			isVisited := make([][]bool, size)
			for y := 0; y < size; y++ {
				isVisited[y] = make([]bool, size)
				for x := 0; x < size; x++ {
					isVisited[y][x] = false
				}
			}

			queue = append(queue, getPos(j, i, size))
			for {
				if len(queue) == 0 {
					break
				}

				first := queue[0]
				queue = queue[1:]

				
			}
		}
	}
}

func getPos(x int, y int, size int) int {
	return (y * size) + x
}
