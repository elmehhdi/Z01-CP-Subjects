package main

import "fmt"

func WeAreUnique(str1, str2 string) int {

	if str1 == "" && str2 == "" {
		return -1
	}
	count := 0
	for i := 0; i < len(str1); i++ {
		isunique := true
		for j := 0; j < len(str2); j++ {
			if str1[i] == str2[j] {
				isunique = false
				break
			}

		}
		if isunique {
			count++
		}

	}
	for i := 0; i < len(str2); i++ {
		isunique := true
		for j := 0; j < len(str1); j++ {
			if str2[i] == str1[j] {
				isunique = false
				break
			}

		}
		if isunique {
			count++
		}

	}
	return count

}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
	fmt.Println(WeAreUnique("hi", "hi"))

}
