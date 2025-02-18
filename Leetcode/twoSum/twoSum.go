package main

import "fmt"

func twoSumOpt(a []int, t int) []int {
	aMap := make(map[int]int)
	for i := range a {
		k := t - a[i]
		if v, ok := aMap[k]; ok {
			return []int{i, v}
		}
		aMap[a[i]] = i
	}
	return []int{-1, -1}
}

func twoSum(a []int, t int) (int, int) {
	for i := 0; i < len(a); i++ {
		k := t - a[i]
		for j := i + 1; j < len(a); j++ {
			if a[j] == k {
				return i, j
			}
		}
	}
	return -1, -1
}

func main() {
	a := []int{2, 4, 6, 8, 9}
	t := 15
	fmt.Println(twoSumOpt(a, t))
}
