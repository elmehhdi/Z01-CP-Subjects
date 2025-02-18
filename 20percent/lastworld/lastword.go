package main

import "fmt"

func LastWord(s string) string {
	if s == "" {
		return "\n"
	}

	word := ""
	i := len(s) - 1

	if s[i] == ' ' {
		for i >= 0 && s[i] == ' ' {
			i--
		}
	}

	for i >= 0 {
		if s[i] == ' ' {
			if word != "" {
				break
			}
		} else {
			word = string(s[i]) + word
		}
		i--
	}

	return word + "\n"
}

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}
