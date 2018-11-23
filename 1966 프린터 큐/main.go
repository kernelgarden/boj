package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var caseNum int

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &caseNum)

	for i := 0; i < caseNum; i++ {
		var n, m int
		fmt.Fscan(reader, &n)
		fmt.Fscan(reader, &m)
		queue := make([]int, n, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &queue[j])
		}

		result := Simulate(queue, m, 1)
		fmt.Println(result)
	}
}

func Simulate(queue []int, targetPos int, curStep int) int {
	sliceLen := len(queue)

	if FindGreaterThen(queue, 0) {
		first := queue[0]
		queue = queue[1 : sliceLen]
		queue = append(queue, first)

		// 시뮬레이션 대상이 옮겨진 경우만 체크
		if targetPos == 0 {
			targetPos = sliceLen - 1
		} else {
			targetPos--
		}

		return Simulate(queue, targetPos, curStep)
	} else {
		if targetPos == 0 {
			return curStep
		} else {
			queue = queue [1 : sliceLen]
			return Simulate(queue, targetPos - 1, curStep + 1)
		}
	}
}

func FindGreaterThen(queue []int, comparePos int) bool {
	target := queue[comparePos]

	for i := 0; i < len(queue); i++ {
		if comparePos == i {
			continue
		}

		if queue[i] > target {
			return true
		}
	}

	return false
}
