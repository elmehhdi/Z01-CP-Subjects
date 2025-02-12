package main

import "fmt"

func CheckNumber(arg string) bool {
	for _, num := range arg {
		if num >= '0' && num <= '9' {
			return true
		}

	}
	return false

}

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
	fmt.Println(CheckNumber("5"))

}
