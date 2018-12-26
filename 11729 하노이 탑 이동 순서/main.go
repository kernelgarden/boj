package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d", &n)

	num, moveInfos := move(n, 1, 2, 3)
	fmt.Println(num)
	for _, moveInfo := range moveInfos {
		fmt.Printf("%d %d\n", moveInfo.src, moveInfo.dest)
	}
}

// 재귀적인 의미로 접근하면 쉽다.
// src, buf, dest는 각각 옮겨올 인덱스, 버퍼 인덱스, 올겨질 대상 인덱스다.
func move(n, src, buf, dest int) (int, []*MoveInfo) {
	acc := 0
	infos := make([]*MoveInfo, 0)

	if n == 1 {
		return 1, []*MoveInfo{MakeMoveInfo(src, dest)}
	}

	// n - 1개를 버퍼로 옮겨둔다.
	num, moveInfos := move(n - 1, src, dest, buf)
	acc += num
	infos = append(infos, moveInfos...)

	// 소스의 마지막 1개를 최종 인덱스로 옮긴다.
	num, moveInfos = move(1, src, buf, dest)
	acc += num
	infos = append(infos, moveInfos...)

	// n - 1개를 최종 인덱스로 옮긴다.
	num, moveInfos = move(n - 1, buf, src, dest)
	acc += num
	infos = append(infos, moveInfos...)

	return acc, infos
}

type MoveInfo struct {
	src, dest int
}

func MakeMoveInfo(src, dest int) *MoveInfo {
	return &MoveInfo{src: src, dest: dest}
}
