package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1200000000000000)
	b := big.NewInt(2000000000000000)

	sum := new(big.Int).Add(a, b)
	sub := new(big.Int).Sub(a, b)
	mul := new(big.Int).Mul(a, b)
	div := new(big.Int).Div(a, b)

	fmt.Println("a + b =", sum)
	fmt.Println("a - b =", sub)
	fmt.Println("a * b =", mul)
	fmt.Println("a / b =", div)
}
