package main

import "fmt"

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println(subarray(arr, 0, 0))
}

func subarray(arr []int, start int, end int) []int {
	if end == len(arr) {
		return arr
	} else if start > end {
		return subarray(arr, 0, end+1)

	} else {
		fmt.Println(arr[start : end+1])
		return subarray(arr, start+1, end)
	}
}
