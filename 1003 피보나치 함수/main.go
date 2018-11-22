package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var cnt int

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &cnt)

	for a := 0; a < cnt; a++ {
		var n int
		fmt.Fscan(reader, &n)

		cache := make([][]int, n + 1, n + 1)

		for i := 0; i <= n; i++ {
			cache[i] = make([]int, 2, 2)
			for j := 0; j < 2; j++ {
				cache[i][j] = -1
			}
		}

		FiboDiv(n, cache)
		fmt.Printf("%d %d\n", cache[n][0], cache[n][1])
	}
}

func FiboDiv(n int, cache [][]int)  {
	if n == 0 {
		cache[0][0] = 1
		cache[0][1] = 0
		return
	} else if n == 1 {
		cache[1][0] = 0
		cache[1][1] = 1
		return
	}

	if cache[n - 2][0] == -1 {
		FiboDiv(n - 2, cache)
	}
	if cache[n - 1][0] == -1 {
		FiboDiv(n - 1, cache)
	}

	cache[n][0] = cache[n - 1][0] + cache[n - 2][0]
	cache[n][1] = cache[n - 1][1] + cache[n - 2][1]
}