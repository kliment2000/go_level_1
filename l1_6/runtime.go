package main

import (
	"fmt"
	"runtime"
	"time"
)

func worker4() {
	fmt.Println("worker4 начал работу")
	time.Sleep(1 * time.Second)
	fmt.Println("worker4 завершен (Goexit)")
	runtime.Goexit()
}

func main() {
	go worker4()
	time.Sleep(2 * time.Second)
}
