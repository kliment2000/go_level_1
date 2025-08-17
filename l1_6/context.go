package main

import (
	"context"
	"fmt"
	"time"
)

func worker3(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker3 завершен по context")
			return
		default:
			fmt.Println("worker3 работает")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go worker3(ctx)

	time.Sleep(3 * time.Second)
}
