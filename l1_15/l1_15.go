package main

import (
	"fmt"
	"runtime"
	"strings"
)

var justString string

func createHugeString(size int) string {
	return string(make([]byte, size))
}

func someFuncLeak() {
	v := createHugeString(100 << 20)
	justString = v[:100]
	fmt.Println("Вызывается функция")
}

func someFuncFixed() {
	v := createHugeString(100 << 20)
	justString = strings.Clone(v[:100])
	fmt.Println("Вызывается функция")
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Занято %v MB\n", m.Alloc/1024/1024)
}

func main() {
	fmt.Println("В изначальном варианте происходит утечка")
	printMemUsage()
	someFuncLeak()
	runtime.GC()
	printMemUsage()

	justString = ""
	runtime.GC()

	fmt.Println("\nКорректный пример реализации")
	printMemUsage()
	someFuncFixed()
	runtime.GC()
	printMemUsage()

	_ = justString // чтоб GoLand не ругался
}
