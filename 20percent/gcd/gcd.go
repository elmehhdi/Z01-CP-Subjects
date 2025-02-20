package main

import "fmt"

func Gcd(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}

	var divisors_of_a []uint
	for i := uint(1); i < a; i++ {
		if a%i == 0 {
			divisors_of_a = append(divisors_of_a, i)

		}
	}

	var divisors_of_b []uint
	for i := uint(1); i < b; i++ {
		if b%i == 0 {
			divisors_of_b = append(divisors_of_b, i)

		}
	}

	gcd := uint(1)
	for _, a_divisor := range divisors_of_a {
		for _, b_divisor := range divisors_of_b {
			if a_divisor == b_divisor && a_divisor > gcd {
				gcd = a_divisor
			}

		}

	}
	return gcd

}

func main() {
	fmt.Println(Gcd(42, 10))
	fmt.Println(Gcd(42, 12))
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))
}
