package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var directions = [4][2]int {
	{0, -1},
	{-1, 0}, {1, 0},
	{0, 1},
}

var scanner = bufio.NewScanner(os.Stdin)

func nextInt() int {
	scanner.Scan()
	r := 0
	for _, c := range scanner.Bytes() {
		r *= 10
		r += int(c - '0')
	}

	return r
}

func nextInts() []int {
	scanner.Scan()

	ints := make([]int, 0)
	for _, c := range scanner.Bytes() {
		ints = append(ints, int(c - '0'))
	}

	return ints
}

func main() {
	scanner.Split(bufio.ScanWords)

	size := nextInt()
	m := make([][]int, size, size)

	from := make([]int, size * size)
	to := make(map[int]map[int]bool)

	for h := 0; h < size; h++ {
		m[h] = make([]int, size, size)

		inputs := nextInts()

		for w := 0; w < size; w++ {
			m[h][w] = inputs[w]

			pos := (h * size) + w
			from[pos] = pos

			if m[h][w] == 1 {
				to[pos] = map[int]bool{pos: true}
			}
		}
	}

	for h := 0; h < size; h++ {
		for w := 0; w < size; w++ {
			pos := (h * size) + w
			connectedPoses := GetConnectedPoses(m, w, h, size)

			for _, connectedPos := range connectedPoses {
				a, b := from[connectedPos], from[pos]

				// 이미 연결되어 있는 노드는 체크하지 않는다.
				if a == b {
					continue
				}

				if len(to[a]) < len(to[b]) {
					a, b = b, a
				}

				// 작은 놈에 링크 되어 있는 애들 큰 놈으로 링크 (본인 포함)
				for v := range to[b] {
					from[v] = a
					to[a][v] = true
				}

				// 이전 링크 정보 삭제
				delete(to, b)
			}
		}
	}

	result := make([]int, 0)
	for _, v := range to {
		result = append(result, len(v))
	}

	sort.Ints(result)
	fmt.Println(len(result))
	for _, v := range result {
		fmt.Println(v)
	}
}

func GetConnectedPoses(m [][]int, x int, y int, mapSize int) []int {
	connectedPoses := make([]int, 0)

	for _, relPos := range directions {
		checkX := x + relPos[0]
		checkY := y + relPos[1]

		if checkX >= 0 && checkX < mapSize && checkY >= 0 && checkY < mapSize {
			if m[checkY][checkX] == 1 {
				connectedPoses = append(connectedPoses, (checkY * mapSize) + checkX)
			}
		}
	}

	return connectedPoses
}
