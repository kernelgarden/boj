package main

/*
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
	var n, m int

	n = nextInt()
	have := make(map[int]bool)
	for i := 0; i < n; i++ {
		have[nextInt()] = true
	}

	m = nextInt()
	for i := 0; i < m; i++ {
		num := nextInt()
		if found, _ := have[num]; found {
			fmt.Fprintf(out, "%d ", 1)
		} else {
			fmt.Fprintf(out, "%d ", 0)
		}
	}

	fmt.Fprintln(out)
	out.Flush()
}
*/
