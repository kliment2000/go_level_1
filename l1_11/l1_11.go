package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			if num1 == num2 {
				result = append(result, num1)
			}
		}
	}
	return result
}

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	c := intersect(a, b)
	fmt.Println(c)
}
