package main

import (
	"fmt"
	"time"
)

func main() {
	N := 5 // время работы программы в секундах
	ch := make(chan int)
	done := make(chan struct{})

	go func() {
		counter := 1
		for {
			select {
			case <-done:
				fmt.Println("Писатель завершает работу")
				return
			case ch <- counter:
				fmt.Println("Отправляю:", counter)
				counter++
				time.Sleep(time.Second)
			}
		}
	}()

	go func() {
		for val := range ch {
			fmt.Println("Прочитал:", val)
		}
	}()

	<-time.After(time.Duration(N) * time.Second)

	fmt.Println("Время вышло, завершаем программу...")
	close(done)
	close(ch)

	<-time.After(2 * time.Second)
}
