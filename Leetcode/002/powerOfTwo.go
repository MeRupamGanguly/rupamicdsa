package main

import "fmt"

func powerOfTwo(n int) bool {
	if n < 1 { //power of 2 will be never negetive.
		return false
	}
	for ; n%2 == 0; n = n / 2 {

	}
	return n == 1
}

func main() {
	fmt.Println(powerOfTwo(46))
}
