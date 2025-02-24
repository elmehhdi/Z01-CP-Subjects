package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	str := os.Args[1]
	searched_char := os.Args[2]
	new_char := os.Args[3]

	if len(searched_char) != 1 || len(new_char) != 1 {
		return
	}

	str_rune := []rune(str)
	searched_char_rune := rune(searched_char[0])
	new_char_rune := rune(new_char[0])

	for i := 0; i < len(str_rune); i++ {
		if str_rune[i] == searched_char_rune {
			str_rune[i] = new_char_rune
		}
	}

	fmt.Println(string(str_rune))
}
