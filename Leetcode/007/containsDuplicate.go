package main

import "fmt"

func containsDuplicate(nums []int) bool {
	myMap := make(map[int]int)
	for i := range nums {
		val, ok := myMap[nums[i]]
		if !ok {
			myMap[nums[i]] = val + 1
		} else {
			return true
		}
	}
	return false
}
func main() {
	nums := []int{1, 2, 4, 6, 8, 4, 7, 9}
	fmt.Println(containsDuplicate(nums))
}
