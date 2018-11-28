package main

import (
	"bufio"
	"fmt"
	"os"
)

var near = [4][2]int {
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0},
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		var w, h int
		fmt.Fscan(reader, &w)
		fmt.Fscan(reader, &h)

		if w == 0 && h == 0 {
			break
		}

		landMap := make([][]int, h, h)
		from := make([]int, h * w, h * w)
		to := make(map[int]map[int]bool)

		for i := 0; i < h; i++ {
			landMap[i] = make([]int, w, w)

			for j := 0; j < w; j++ {
				curPos := (i * w) + j

				from[curPos] = curPos
				fmt.Fscan(reader, &landMap[i][j])

				// 현재 위치가 섬이 아니면 체크할 필요도 없음
				if landMap[i][j] == 0 {
					continue
				}

				to[curPos] = map[int]bool{curPos: true}

				connectedLands := CheckLand(landMap, w, h, j, i)

				for k := 0; k < len(connectedLands); k++ {
					a, b := from[curPos], from[connectedLands[k]]

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
			}
		}

		fmt.Println(len(to))
	}
}

func CheckLand(land [][]int, xSize int, ySize, x int, y int) []int {
	connectedLands := make([]int, 0)

	for _, pos := range near {
		checkX := x + pos[0]
		checkY := y + pos[1]

		if (checkX >= 0) && (checkX < xSize) && (checkY >= 0) && (checkY < ySize) {
			if land[checkY][checkX] == 1 {
				connectedLands = append(connectedLands, (checkY * xSize) + checkX)
			}
		}
	}

	return connectedLands
}

