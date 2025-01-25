package main

import "fmt"

func palindrme(x int) bool {
	if x < 0 {
		return false
	}
	reverseNumber := 0
	for number := x; number > 0; number = number / 10 {
		reverseNumber = reverseNumber*10 + number%10
	}
	return reverseNumber == x
}

func main() {
	fmt.Println(palindrme(7121))
}
