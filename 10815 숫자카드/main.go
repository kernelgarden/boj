package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var n, m int
	n = nextInt()
	have := make([]int, n)
	for i := 0; i < n; i++ {
		have[i] = nextInt()
	}

	sort.Sort(sort.IntSlice(have))

	m = nextInt()
	for i := 0; i < m; i++ {
		target := nextInt()
		if found, _ := binSearch(have, target); found {
			fmt.Fprintf(out, "%d ", 1)
		} else {
			fmt.Fprintf(out, "%d ", 0)
		}
	}
	fmt.Fprintln(out)
	out.Flush()
}

func binSearch(s []int, target int) (bool, int) {
	found := false
	foundIdx := -1
	size := len(s)
	var mid int
	start, end := 0, size - 1

	for start <= end {
		mid = (start + end) / 2

		if s[mid] == target {
			found = true
			foundIdx = mid
			break
		}

		if target > s[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return found, foundIdx
}
