package main

import (
	"fmt"
)

func reverse(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

func main() {
	input := "Hello главрыба 👋"
	fmt.Println("Исходная строка:", input)

	fmt.Println("Перевёрнутая строка:", reverse(input))
}
