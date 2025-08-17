package main

import (
	"fmt"
	"time"
)

func worker1() {
	for i := 1; i <= 5; i++ {
		fmt.Println("worker1:", i)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("worker1 завершен по условию")
}

func main() {
	go worker1()
	time.Sleep(3 * time.Second)
}
