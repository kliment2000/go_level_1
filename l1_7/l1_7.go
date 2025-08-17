package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{m: make(map[string]int)}
}

func (s *SafeMap) Inc(key string) {
	s.mu.Lock()
	s.m[key]++
	s.mu.Unlock()
}

func (s *SafeMap) Get(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.m[key]
}

func main() {
	const workers = 10
	const perWorker = 10000

	sm := NewSafeMap()
	var wg sync.WaitGroup

	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := 0; i < perWorker; i++ {
				sm.Inc("counter")
			}
		}(w)
	}

	wg.Wait()
	want := workers * perWorker
	got := sm.Get("counter")
	fmt.Printf("OK=%v", got == want)
}
