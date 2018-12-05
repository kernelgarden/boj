package main

import (
	"bufio"
	"fmt"
	"math"
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

	/*
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%v ", nodes[i][j])
		}
		fmt.Println()
	}
	*/
	shortestDist := math.MaxInt32
	findShortestDist(nodes, size, &shortestDist)
	fmt.Println(shortestDist)
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

/*
func findShortestDistUtil(nodes [][]int, size int, startX int, startY int, startGroupId int, depth int, shortestDist *int) {
	// 이 경우는 탐색할 필요가 이미 없다.
	if depth >= *shortestDist {
		return
	}

	connectedLands, connectedSeas := GetConnectedNodes(nodes, size, getPos(startX, startY, size))

	// 인접 육지를 먼저 체크한다.
	for _, connectedLand := range connectedLands {
		x, y := getXY(connectedLand, size)
		if nodes[y][x] != startGroupId {
			if depth < *shortestDist {
				*shortestDist = depth
			}
		}
	}

	for _, connectedSea := range connectedSeas {
		x, y := getXY(connectedSea, size)
		findShortestDistUtil(nodes, size, x, y, startGroupId, depth + 1, shortestDist)
	}
}

func findShortestDist(nodes [][]int, size int) int {
	shortestDist := math.MaxInt32
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			// 바다인 곳은 탐색 시작 노드가 될 수 없다.
			if nodes[i][j] == 0 {
				continue
			}

			findShortestDistUtil(nodes, size, j, i, nodes[i][j], 1, &shortestDist)
		}
	}

	return shortestDist
}
*/

type PosDepthPair struct {
	pos int
	depth int
}

func NewPosDepthPair(pos, depth int) PosDepthPair {
	return PosDepthPair{ pos, depth}
}

func findShortestDist(nodes [][]int, size int, shortestDist *int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			// 탐색 시작 노드가 바다면 패스
			if nodes[i][j] == 0 {
				continue
			}

			startGroupId := nodes[i][j]

			isFind := false
			queue := make([]PosDepthPair, 0)
			isVisited := make([][]bool, size)
			for y := 0; y < size; y++ {
				isVisited[y] = make([]bool, size)
				for x := 0; x < size; x++ {
					isVisited[y][x] = false
				}
			}

			queue = append(queue,NewPosDepthPair(getPos(j, i, size), 0))
			for {
				if len(queue) == 0 || isFind {
					break
				}

				target := queue[0]
				queue = queue[1:]

				// 이미 방문한 노드면 재 방문 필요가 없음
				if x, y := getXY(target.pos, size); isVisited[y][x] == true {
					continue
				} else {
					isVisited[y][x] = true
				}

				// 최단 노드가 아니라면 탐색 포기
				if target.depth >= *shortestDist {
					break
				}

				connectedLands, connectedSeas := GetConnectedNodes(nodes, size, target.pos)

				// 인접한 육지부터 체크
				for _, connectedLand := range connectedLands {
					x, y := getXY(connectedLand, size)
					// 다른 최단 노드를 찾은 경우다.
					if nodes[y][x] != startGroupId {
						isFind = true
						if target.depth < *shortestDist {
							*shortestDist = target.depth
						}
						break
					}
				}

				for _, connectedSea := range connectedSeas {
					queue = append(queue, NewPosDepthPair(connectedSea, target.depth + 1))
				}
			}
		}
	}
}

func getPos(x int, y int, size int) int {
	return (y * size) + x
}

func getXY(pos int, size int) (int, int) {
	return pos % size, pos / size
}
