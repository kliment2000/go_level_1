package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер %d завершает работу\n", id)
			return
		case job := <-jobs:
			fmt.Printf("Воркер %d получил: %d\n", id, job)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование: go run main.go <количество_воркеров>")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println("Ошибка: количество воркеров должно быть положительным числом")
		return
	}

	jobs := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println("\nПолучен сигнал прерывания, закрываю контекст")
		cancel()
	}()

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	counter := 1
loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		default:
			jobs <- counter
			counter++
			time.Sleep(300 * time.Millisecond)
		}
	}

	wg.Wait()
	fmt.Println("Главная горутина завершена")
}
