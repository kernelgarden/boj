package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int

	fmt.Fscanf(reader, "%d\n", &n)
	have := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d", &have[i])
	}

	fmt.Fscanf(reader, "\n%d\n", &m)
	search := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d", &search[i])
	}

	sort.Sort(sort.IntSlice(have))

	result := make([]int, 0, m)
	for _, num := range search {
		found, _ := binSearch(have, num)

		if found {
			result = append(result, 1)
		} else {
			result = append(result, 0)
		}
	}

	for _, num := range result {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
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
