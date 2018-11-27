package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, v int

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d", &n, &v)

	// 연결된 요소의 대표 노드를 값으로 가지고 있는 해시맵
	from := make(map[int]int)
	// 대표 노드의 경우만 이 해시맵에 들어 있다.
	// 여기에 들어 있는 값은 링크 되어 있는 노드들이다.
	to := make(map[int]map[int]bool)

	for i := 1; i <= n; i++ {
		from[i] = i
		to[i] = map[int]bool{i: true}
	}

	for i := 0; i < v; i++ {
		var n1, n2 int
		fmt.Fscan(reader, &n1)
		fmt.Fscan(reader, &n2)
		a, b := from[n1], from[n2]

		// 이미 같은 요소에 연결되 있는 상태다.
		if a == b {
			continue
		}

		// 링크되어 있는 노드가 작은 요소를 합치는 대상으로 본다.
		if len(to[a]) < len(to[b]) {
			a, b = b, a
		}

		// 작은 요소의 링크들을 큰 요소에 합친다.
		// (대표 노드의 경우 to 해시에 초기화 시점에 본인의 값을 넣어줬다.)
		for link := range to[b] {
			from[link] = a
			to[a][link] = true
		}

		// 작은 요소는 삭제한다.
		delete(to, b)
	}

	fmt.Print(len(to))
}