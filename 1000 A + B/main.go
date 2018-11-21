package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b int

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d", &a, &b)
	fmt.Printf("%d", a + b)
}