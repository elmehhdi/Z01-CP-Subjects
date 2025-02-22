package main

import "fmt"

func RepeatAlpha(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		char := s[i]
		repeat := 1

		if char >= 'a' && char <= 'z' {
			repeat = int(char - 'a' + 1)
		}

		if char >= 'A' && char <= 'Z' {
			repeat = int(char - 'A' + 1)
		}

		for j := 0; j < repeat; j++ {
			result += string(char)
		}

		if !(char >= 'a' && char <= 'z') && !(char >= 'A' && char <= 'Z') {
			result += string(char)
		}
	}

	return result
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
