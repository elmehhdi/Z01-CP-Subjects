package main

import "fmt"

func FromTo(from int, to int) string {
	result := ""
	if from > 99 || from < 0 || to > 99 || to < 0 {
		return "Invalid\n"
	}
	if from <= to {
		for i := from; i <= to; i++ {
			if i != from {
				result += ", "
			}
			result += fmt.Sprintf("%02d", i)

		}
	} else if from > to {
		for i := from; i >= to; i-- {
			if i != from {
				result += ", "
			}
			result += fmt.Sprintf("%02d", i)
		}

	}

	return result + "\n"
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
	fmt.Print(FromTo(4, 26))
	fmt.Print(FromTo(200, 6))
	fmt.Print(FromTo(56, 46))
	fmt.Print(FromTo(74, 87))

}
