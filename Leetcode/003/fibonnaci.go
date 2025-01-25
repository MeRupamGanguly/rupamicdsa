package main

import "fmt"

func fibonnaci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	f := 0
	s := 1
	t := 0
	for i := 2; i <= n; i++ {
		t = f + s
		f = s
		s = t
	}
	return t
}

func main() {
	fmt.Println(fibonnaci(7))
}
