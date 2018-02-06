package main

import (
	"fmt"
	"os"
)

func f(n int) int {
	fmt.Fprintf(os.Stdout, "%d ", n)
	if n == 0 || n == 1 {
		return n
	}
	if n%2 == 0 {
		return f(n / 2)
	}
	return f(3*n + 1)
}

func main() {
	var n int
	for true {
		c, err := fmt.Fscanf(os.Stdin, "%d", &n)
		if c == 0 || err != nil {
			break
		}
		f(n)
		fmt.Fprintf(os.Stdout, "\n")
	}
}
