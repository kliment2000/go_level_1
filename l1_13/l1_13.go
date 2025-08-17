package main

import "fmt"

func main() {
	a := 5
	b := 11
	fmt.Println(a, b)
	b = b - a
	a = a + b
	b = a - b
	fmt.Println(a, b)

	fmt.Println()

	c := 20
	d := 13
	fmt.Println(c, d)
	c = c ^ d
	d = c ^ d
	c = c ^ d
	fmt.Println(c, d)
}
