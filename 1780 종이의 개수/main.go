package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanf(reader, "%d", &n)

	paper := make([][]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "\n")
		paper[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscanf(reader, "%d", &paper[i][j])
		}
	}

	a, b, c := GetPapersNum(paper, MakeBound(0, 0, n))
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func GetPapersNum(paper [][]int, bound Bound) (int, int, int) {
	if bound.size == 1 {
		return ConvertFormat(paper[bound.y][bound.x], 1)
	}

	if CheckAllSame(paper, bound) {
		return ConvertFormat(paper[bound.y][bound.x], 1)
	} else {
		a, b, c := 0, 0, 0
		newBounds := SliceBound(bound)
		for i := 0; i < 9; i++ {
			sa, sb, sc := GetPapersNum(paper, newBounds[i])
			a += sa
			b += sb
			c += sc
		}

		return a, b, c
	}
}

func CheckAllSame(paper [][]int, bound Bound) bool {
	num := paper[bound.y][bound.x]

	for i := 0; i < bound.size; i++ {
		for j := 0; j < bound.size; j++ {
			if paper[bound.y + i][bound.x + j] != num {
				return false
			}
		}
	}

	return true
}

func ConvertFormat(num int, n int) (int, int, int) {
	switch num {
	case -1:
		return n * n, 0, 0
	case 0:
		return 0, n * n, 0
	case 1:
		return 0, 0, n * n
	}

	return 0, 0, 0
}

func SliceBound(bound Bound) []Bound {
	nextN := bound.size / 3
	newBounds := make([]Bound, 9)

	for i := 0; i < 9; i++ {
		x := bound.x + (nextN * (i % 3))
		y := bound.y + (nextN * (i / 3))
		newBounds[i] = MakeBound(x, y, nextN)
	}

	return newBounds
}

type Bound struct {
	x, y int
	size int
}

func MakeBound(x, y, size int) Bound {
	return Bound{x: x, y: y, size: size}
}
