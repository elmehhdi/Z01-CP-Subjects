package main

import "fmt"

func isprime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false

		}

	}
	return true

}

func FindPrevPrime(nb int) int {

	for i := nb; i >= 2; i-- {
		if isprime(i) {
			return i
		}

	}
	return 0

}

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
	fmt.Println(FindPrevPrime(20))
	fmt.Println(FindPrevPrime(32))
	fmt.Println(FindPrevPrime(90))
	fmt.Println(FindPrevPrime(-5))
	fmt.Println(FindPrevPrime(1))

}
