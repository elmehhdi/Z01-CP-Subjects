package main

import "fmt"

func ConcatAlternate(slice1, slice2 []int) []int {
	newslice := []int{}

	if len(slice1) < len(slice2) {
		slice1, slice2 = slice2, slice1
	}

	i := 0
	for i < len(slice2) {
		newslice = append(newslice, slice1[i], slice2[i])
		i++
	}

	newslice = append(newslice, slice1[i:]...)

	return newslice
}

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}
