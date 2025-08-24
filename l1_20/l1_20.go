package main

import (
	"fmt"
)

func reverse(runes []rune, i, j int) {
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
}

func reverseWords(s string) string {
	runes := []rune(s)

	reverse(runes, 0, len(runes)-1)

	start := 0
	for i := 0; i <= len(runes); i++ {
		if i == len(runes) || runes[i] == ' ' {
			reverse(runes, start, i-1)
			start = i + 1
		}
	}

	return string(runes)
}

func main() {
	input := "snow dog sun"

	fmt.Println("Результат:", reverseWords(input))
}
