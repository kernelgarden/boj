package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	red = 0
	blue = 1
	green = 2
)

var scanner = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func nextInt() int {
	scanner.Scan()
	r := 0
	f := 1
	for _, c := range scanner.Bytes() {
		if c == '-' {
			f = -1
			continue
		}
		r *= 10
		r += int(c - '0')
	}
	return r * f
}

func main() {
	scanner.Split(bufio.ScanWords)

	n := nextInt()
	costs := make([][]int, n)
	minCache := make([][]int, n)

	for i := 0; i < n; i++ {
		costs[i] = make([]int, 3)
		minCache[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			costs[i][j] = nextInt()
			minCache[i][j] = -1
		}
	}

	red := CalcCost(0, red, n, costs, minCache)
	blue := CalcCost(0, blue, n, costs, minCache)
	green := CalcCost(0, green, n, costs, minCache)

	fmt.Fprintln(out, min(red, blue, green))
	out.Flush()
}

func CalcCost(depth int, cur int, n int, cost [][]int, minCache [][]int) int {
	if depth  == n {
		return 0
	}

	if minCache[depth][cur] != -1 {
		return minCache[depth][cur]
	}

	var next1, next2 int
	switch cur {
	case red:
		next1, next2 = blue, green
		break
	case blue:
		next1, next2 = red, green
		break
	case green:
		next1, next2 = red, blue
		break
	}

	sum1 := CalcCost(depth + 1, next1, n, cost, minCache)
	sum2 := CalcCost(depth + 1, next2, n, cost, minCache)

	var min int
	if sum1 <= sum2 {
		min = cost[depth][cur] + sum1
		minCache[depth][cur] = min
	} else {
		min = cost[depth][cur] + sum2
		minCache[depth][cur] = min
	}

	return min
}

func min(nums ...int) int {
	min := math.MaxInt32
	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min
}
