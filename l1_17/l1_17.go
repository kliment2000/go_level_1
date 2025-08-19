package main

import "fmt"

func binSearch(a []int, key int) int {
	n := len(a)
	left, right := -1, n
	for left+1 < right {
		mid := (left + right) / 2
		if a[mid] == key {
			return mid
		}
		if a[mid] > key {
			right = mid
		} else {
			left = mid
		}
	}
	return -1
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 8, 10, 12, 15, 16, 18, 20}

	fmt.Println(binSearch(a, -1))
	fmt.Println(binSearch(a, 11))
	fmt.Println(binSearch(a, 12))
	fmt.Println(binSearch(a, 30))
}
