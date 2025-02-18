package main

import "fmt"

func waterBottles(n, x int) int {
	res := n
	rem := n % x
	for {
		n = n / x
		if n == 0 {
			break
		}
		res += n
		n = rem + n
		rem = n % x
	}
	return res
}

func main() {
	fmt.Println(waterBottles(15, 4))
}
