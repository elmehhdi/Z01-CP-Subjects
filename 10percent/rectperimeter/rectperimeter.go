package main

import (
	"fmt"
)

func RectPerimeter(w, h int) int {
	if w <= 0 || h <= 0 {
		return -1
	}

	perimeter := 2 * (w + h)

	return perimeter
}

func main() {
	fmt.Println(RectPerimeter(10, 5))
	fmt.Println(RectPerimeter(15, 6))
	fmt.Println(RectPerimeter(4, 0))
}
