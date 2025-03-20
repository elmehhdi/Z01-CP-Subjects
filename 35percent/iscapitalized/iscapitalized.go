package main

import "fmt"

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}

	i := 0
	n := len(s)

	for i < n {
		for i < n && s[i] == ' ' {
			i++
		}

		if i >= n {
			break
		}

		firstChar := s[i]

		if firstChar >= 'a' && firstChar <= 'z' {
			return false
		}

		for i < n && s[i] != ' ' {
			i++
		}
	}

	return true
}

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}
