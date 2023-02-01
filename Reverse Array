package main

import "fmt"

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println(reverse(arr))
}

func reverse(arr []int) []int {
	start := 0
	end := len(arr) - 1
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
	return arr
}
