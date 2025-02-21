package main

import "fmt"

func HashCode(dec string) string {

	hash := ""
	length := len(dec)

	for _, char := range dec {
		newchar := (int(char) + length) % 127
		if newchar < 32 {
			newchar = newchar + 33
		}
		hash += string(newchar)

	}
	return hash

}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
