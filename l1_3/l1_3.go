package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func worker(id int, ch chan int) {
	for {
		job := <-ch
		fmt.Printf("Воркер %d получил: %d\n", id, job)
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
	for i := 1; i <= n; i++ {
		go worker(i, jobs)
	}

	counter := 1
	for {
		jobs <- counter
		counter++
		time.Sleep(time.Second)
	}
}
