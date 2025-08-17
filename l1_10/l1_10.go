package main

import (
	"fmt"
	"math"
)

func main() {
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := make(map[int][]float64)

	for _, t := range temp {
		var key int
		if t > 0 {
			key = int(math.Floor(t/10) * 10)
		} else {
			key = int(math.Ceil(t/10) * 10)
		}
		m[key] = append(m[key], t)
	}

	for key, value := range m {
		fmt.Printf("%v: %v\n", key, value)
	}
}
