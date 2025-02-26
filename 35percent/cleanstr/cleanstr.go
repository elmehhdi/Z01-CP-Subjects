package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	str := os.Args[1]
	result := ""

	space := false
	started := false

	for i := 0; i < len(str); i++ {
		char := str[i]

		if char == ' ' || char == '\t' {
			if started {
				space = true
			}
		} else {
			if space {
				result += " "
			}
			result += string(char)
			space = false
			started = true
		}
	}

	fmt.Println(result)
}
