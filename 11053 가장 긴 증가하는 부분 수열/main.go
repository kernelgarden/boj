package main

import (
	"bufio"
	"fmt"
	"os"
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
	nums := make([]int, n, n)

	for i := 0; i < n; i++ {
		nums[i] = nextInt()
	}

	fmt.Fprintln(out, getLIS(nums, n))
	out.Flush()
}

func getLIS(nums []int, n int) int {
	longest := 0

	cache := make([]int, n)
	for i := 0; i < n; i++ {
		cache[i] = -1
	}

	for idx := range nums {
		longest = max(getLength(nums, n, idx, cache), longest)
	}

	return longest
}

/*
// 첫번째 시도였다. 모든 케이스를 체크하지 못했다.
func getLength(nums []int, n int, start int, cache []int) int {
	if cache[start] != -1 {
		return cache[cache[start]]
	}

	prev := nums[start]
	length := 1

	for i := start + 1; i < n; i++ {
		if nums[i] > prev {
			cache[i] = start
			prev = nums[i]
			length++
		}
	}

	return length
}
*/

func getLength(num []int, n int, start int, cache []int) int {
	if start >= n {
		return 0
	}

	if cache[start] != -1 {
		return cache[start]
	}

	longest := 0

	for i := start + 1; i < n; i++ {
		if num[i] > num[start] {
			longest = max(getLength(num, n, i, cache), longest)
		}
	}

	cache[start] = longest + 1
	return cache[start]
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}