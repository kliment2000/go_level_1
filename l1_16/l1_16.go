package main

import "fmt"

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	var left []int
	var right []int
	var mid []int

	base := a[len(a)/2]
	for i := range a {
		if a[i] < base {
			left = append(left, a[i])
		} else if a[i] > base {
			right = append(right, a[i])
		} else {
			mid = append(mid, a[i])
		}
	}
	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, mid...), right...)
}

func main() {
	a := []int{5, 9, 7, 8, 6, 10, 1, 2, 4, 3}
	fmt.Println(a)

	a = quickSort(a)
	fmt.Println(a)
}
