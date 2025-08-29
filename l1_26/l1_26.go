package main

import (
	"fmt"
	"unicode"
)

func uniqueSymbols(s string) bool {
	seen := make(map[rune]bool)

	for _, ch := range s {
		lch := unicode.ToLower(ch)
		if seen[lch] {
			return false
		}
		seen[lch] = true
	}
	return true
}

func main() {
	fmt.Println(uniqueSymbols("abcd"))
	fmt.Println(uniqueSymbols("abCdefAaf"))
	fmt.Println(uniqueSymbols("aabcd"))
	fmt.Println(uniqueSymbols("привет"))
}
