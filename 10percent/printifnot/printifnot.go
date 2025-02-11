package main

import (
	"fmt"
)

func PrintIfNot(str string) string {
	if str == "" {
		return "G\n"
	}

	if len(str) < 3 {
		return "G\n"
	}

	return "Invalid Input\n"
}

func main() {
	fmt.Print(PrintIfNot("abcdefzghskkxnjskslnxn"))
	fmt.Print(PrintIfNot("abc"))
	fmt.Print(PrintIfNot(" "))
	fmt.Print(PrintIfNot("61"))
}
