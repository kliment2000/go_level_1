package main

import (
	"fmt"
	"sync"
)

func pow2(x int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := x * x
	fmt.Println(result)
}

func main() {
	mas := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	for _, m := range mas {
		wg.Add(1)
		go pow2(m, &wg)
	}
	wg.Wait()
}
