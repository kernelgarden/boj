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

func nextString() string {
	scanner.Scan()

	s := make([]byte, 0, 10)

	for _, c := range scanner.Bytes() {
		s = append(s, c)
	}

	return string(s[:])
}

func main() {
	scanner.Split(bufio.ScanWords)

	pattern := nextString()
	matchNum := nextInt()
	matchTargets := make([][]rune, 0, matchNum)
	for i := 0; i < matchNum; i++ {
		matchTarget := nextString()
		matchTargets = append(matchTargets, []rune(matchTarget))
	}

	matchedTargets := matchAll([]rune(pattern), matchNum, matchTargets)
	for _, target := range matchedTargets {
		fmt.Fprintln(out, string(target))
	}

	out.Flush()
}

func removeDup(s []rune) []rune {
	sLen := len(s)
	ret := make([]rune, 0, sLen)

	if sLen == 0 {
		return ret
	}

	before := s[0]
	ret = append(ret, before)
	for i := 1; i < sLen; i++ {
		if before == '*' && s[i] == '*' {
			continue
		} else {
			ret = append(ret, s[i])
			before = s[i]
		}
	}

	return ret
}

func matchAll(pattern []rune, num int, matchTargets [][]rune) [][]rune {
	pattern = removeDup(pattern)
	matchedTargets := make([][]rune, 0, num)

	for _, target := range matchTargets {
		if match(pattern, target) {
			matchedTargets = append(matchedTargets, target)
		}
	}

	return matchedTargets
}

func match(pattern []rune, target []rune) bool {
	patternLen := len(pattern)
	targetLen := len(target)

	if patternLen == 0 {
		//fmt.Printf("patternLen: %v, targetLen: %v\n", patternLen, targetLen)
		if targetLen == 0 {
			return true
		} else {
			return false
		}
	}

	if pattern[0] == '*' {
		if patternLen == 1 {
			return true
		}

		next := pattern[1]
		isMatched := false
		lastMatchedIdx := -1
		for idx, c := range target {
			// 매치 성공을 하더라도 최대한 매치를 해본다.
			if c == next {
				isMatched = true
				lastMatchedIdx = idx
			}

			if (patternLen - 2) >= (targetLen - idx) {
				break
			}
		}

		if isMatched {
			return match(pattern[2:], target[lastMatchedIdx + 1:])
		} else {
			return false
		}
	} else {
		if pattern[0] == target[0] {
			return match(pattern[1:], target[1:])
		} else {
			return false
		}
	}
}