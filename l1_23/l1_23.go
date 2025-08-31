package main

import (
	"fmt"
	"runtime"
)

func printMemUsage() {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Printf("Занято %v MB\n", m.Alloc/1024/1024)
}

type Big struct {
	data [1024 * 1024]int
}

func main() {
	s := make([]*Big, 3)
	for i := range s {
		s[i] = &Big{}
	}
	printMemUsage()

	i := 1

	copy(s[i:], s[i+1:])
	s = s[:len(s)-1]
	printMemUsage()
	fmt.Println(s)
}
