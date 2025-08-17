package main

import (
	"fmt"
)

func setBit(n int64, i uint, bit int) int64 {
	if bit == 1 {
		return n | (1 << (i - 1))
	}
	return n &^ (1 << (i - 1))
}

func main() {
	var n int64 = 5
	fmt.Printf("Исходное: %d (%04b)\n", n, n)

	n = setBit(n, 1, 0)
	fmt.Printf("После установки бита 1 в 0: %d (%04b)\n", n, n)

	n = setBit(n, 2, 1)
	fmt.Printf("После установки бита 2 в 1: %d (%04b)\n", n, n)
}
