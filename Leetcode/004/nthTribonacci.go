package main

import "fmt"

func nthTribo(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	f := 0
	s := 1
	t := 1
	r := 0
	for i := 3; i <= n; i++ {
		r = f + s + t
		f = s
		s = t
		t = r
	}
	return r
}

func main() {
	fmt.Println(nthTribo(6))
}
