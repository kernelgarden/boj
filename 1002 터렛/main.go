package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var cnt int

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &cnt)

	for i := 0; i < cnt; i++ {
		var x1, y1, r1 int
		var x2, y2, r2 int

		fmt.Fscan(reader, &x1, &y1, &r1)
		fmt.Fscan(reader, &x2, &y2, &r2)

		fmt.Println(GetAvailableCount(Point{x1, y1, r1}, Point{x2, y2, r2}))
	}
}

type Point struct {
	x int
	y int
	r int
}

func Dist(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p2.x-p1.x), 2) + math.Pow(float64(p2.y-p1.y), 2))
}

func GetAvailableCount(p1, p2 Point) int {
	distOfPoints := Dist(p1, p2)

	if (p1.x == p2.x) && (p1.y == p2.y) {
		// 동심원 케이스
		if p1.r == p2.r {
			return -1
		} else {
			return 0
		}
	} else {
		// 동심원이 아닌 케이스

		rSub := math.Abs(float64(p1.r - p2.r))
		rAdd := float64(p1.r + p2.r)

		if (distOfPoints > rSub) && (distOfPoints < rAdd) {
			return 2
		} else if (distOfPoints == rSub) || (distOfPoints == rAdd) {
			return 1
		} else {
			return 0
		}
	}
}
