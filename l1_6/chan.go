package main

import (
	"fmt"
	"time"
)

func worker2(done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("worker2 завершен по done")
			return
		default:
			fmt.Println("worker2 работает")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	done := make(chan struct{})

	go worker2(done)

	time.Sleep(2 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
}
