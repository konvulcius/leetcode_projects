package main

import "fmt"

func recursive_search(head []int, target, start int) int {
	if head[0] == target {
		return start
	}
	if len(head) == 1 {
		return -1
	}

	max := len(head) - 1
	// исключение для конкретной задачи
	if head[0] < head[max] && (target < head[0] || target > head[max]) {
		return -1
	}

	length := len(head)
	if index := recursive_search(head[0:length/2], target, start); index != -1 {
		return index
	}
	return recursive_search(head[length/2:length], target, start+length/2)
}

// используем binary search в рекурсивном варианте с исключением для конкретного случая
func search(nums []int, target int) int {
	return recursive_search(nums, target, 0)
}

func main() {
	var arr = []int{1, 3}
	fmt.Println(search(arr, 2))
}
