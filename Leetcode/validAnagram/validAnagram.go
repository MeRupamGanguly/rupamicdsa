package main

import (
	"fmt"
)

// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	myMap := make(map[rune]int)
// 	for _, v := range s {
// 		myMap[v]++
// 	}
// 	for _, v := range t {
// 		myMap[v]--
// 		if myMap[v] == 0 {
// 			delete(myMap, v)
// 		}
// 	}
// 	return len(myMap) == 0
// }

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	arr := make(map[int]int)
	for i := 0; i < 26; i++ {
		arr[97+i] = 0
	}
	sRune := []rune(s)
	tRune := []rune(t)
	for i := range sRune {
		arr[int(sRune[i])] += 1
		arr[int(tRune[i])] -= 1
	}

	for i := range arr {
		if arr[i] != 0 {
			return false
		}
	}
	return true

}
func main() {
	fmt.Println(isAnagram("aacc", "aacc"))
	fmt.Println(isAnagram("rat", "tac"))
}
