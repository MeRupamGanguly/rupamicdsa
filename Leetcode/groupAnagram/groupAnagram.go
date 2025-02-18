package main

import "fmt"

func groupAnagram(strs []string) [][]string {
	anagramMap := make(map[[26]int][]string)
	for i, s := range strs {
		var asciiArr [26]int
		for i := 0; i < len(s); i++ {
			fmt.Println("s[i]: ", s[i], "a: ", 'a', "s[i]-'a': ", s[i]-'a')
			asciiArr[s[i]-'a']++
		}
		// need to put asciiArr into memory but not as a array as a string
		anagramMap[asciiArr] = append(anagramMap[asciiArr], strs[i])
		fmt.Println(anagramMap)
	}
	arrOfArr := [][]string{}
	for _, v := range anagramMap {
		arrOfArr = append(arrOfArr, v)
	}
	return arrOfArr
}

func main() {
	fmt.Println(groupAnagram([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagram([]string{"bdddddddddd", "bbbbbbbbbbc"}))
	// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
}
