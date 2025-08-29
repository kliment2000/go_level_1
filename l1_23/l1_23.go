package main

import (
	"fmt"
	"runtime"
	"strings"
)

func printMemUsage() {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Printf("Занято %v MB\n", m.Alloc/1024/1024)
}

func main() {
	huge := strings.Repeat("A", 10*1024*1024)

	s := []string{"Go", "Rust", "Python", huge}
	printMemUsage()

	// Вариант 1: с утечкой
	s1 := s
	copy(s1[2:], s1[3:])
	s1 = s1[:len(s1)-1]
	printMemUsage()

	// Вариант 2: без утечки
	s2 := s
	copy(s2[2:], s2[3:])
	s2[len(s2)-1] = ""
	s2 = s2[:len(s2)-1]
	printMemUsage()
}
