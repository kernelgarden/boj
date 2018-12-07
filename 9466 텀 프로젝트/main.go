package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func nextInt() int {
	scanner.Scan()
	r := 0
	for _, c := range scanner.Bytes() {
		r *= 10
		r += int(c - '0')
	}

	return r
}

func nextInts() []int {
	scanner.Scan()

	ints := make([]int, 0)
	for _, c := range scanner.Bytes() {
		ints = append(ints, int(c - '0'))
	}

	return ints
}

func main() {
	scanner.Split(bufio.ScanWords)
	caseNum := nextInt()

	for c := 0; c < caseNum; c++{
		stuNum := nextInt()
		students := make([]*Student, 0)
		// 학생의 번호가 1부터 시작한다. 코딩의 편의상 더미를 채워 넣는다.
		students = append(students, NewStudent(0, 0))
		for i := 1; i <= stuNum; i++ {
			to := nextInt()
			students = append(students, NewStudent(i, to))
		}
		fmt.Println(stuNum - CalcCycleNum(students, stuNum))
	}
}

type Student struct {
	val int
	to int
}

func NewStudent(val int, to int) *Student {
	return &Student{val: val, to: to}
}

func CalcCycleNum(students []*Student, stuNum int) int {
	totalCycleNum := 0
	isVisited := make([]bool, stuNum + 1)
	for i := 0; i <= stuNum; i++ {
		isVisited[i] = false
	}

	for i := 1; i <= stuNum; i++ {
		if isVisited[i] {
			if students[i].to == students[i].val {
				totalCycleNum++
			}
			continue
		}

		isCycle, cycleNum := Traverse(students, i, isVisited)
		if isCycle {
			totalCycleNum += cycleNum
		}
	}

	return totalCycleNum
}

func Traverse(students []*Student, targetPos int, isVisited []bool) (bool, int) {
	isCycle := false
	cycleNum := 0

	startPos := targetPos

	for {
		if isVisited[targetPos] {
			if targetPos == startPos {
				isCycle = true
			} else {
				cycleNum = 0
			}
			break
		} else {
			isVisited[targetPos] = true
		}

		target := students[targetPos]
		targetPos = target.to
		cycleNum++
	}

	return isCycle, cycleNum
}
