package main

import "fmt"

func main() {
	count := 0
	for a := 9; a >= 2; a-- {
		for b := a - 1; b >= 1; b-- {
			for c := b - 1; c >= 0; c-- {
				if count > 0 {
					fmt.Print(", ")
				}
				fmt.Printf("%d%d%d", a, b, c)
				count++
			}
		}
	}
	fmt.Println()
}
