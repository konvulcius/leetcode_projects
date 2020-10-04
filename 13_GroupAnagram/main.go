package main

import (
	"fmt"
	"sort"
)

func sortWordLetters(str string) string {
	letters := []byte(str)
	ints := make([]int, len(str))

	for i, val := range letters {
		ints[i] = int(val)
	}
	sort.Ints(ints)
	for i, val := range ints {
		letters[i] = byte(val)
	}
	return string(letters)
}

func groupAnagrams(strs []string) [][]string {
	var (
		ans [][]string
		ansMap = make(map[string][]string)
	)

	for _, str := range strs {
		key := sortWordLetters(str)
		ansMap[key] = append(ansMap[key], str)
	}
	for _, str := range ansMap {
		ans = append(ans, str)
	}
	return ans
}

func main() {
	strs := []string{"eat","tea","tan","ate","nat","bat"}
	fmt.Println(groupAnagrams(strs))
	//Output: [["bat"],["nat","tan"],["ate","eat","tea"]
}
