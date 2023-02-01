package main

import "fmt"

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println(rightshift(arr, 2))
}

func rightshift(arr []int, d int) []int {
	var arr1 []int
	for i := len(arr) - d; i < len(arr); i++ {
		arr1 = append(arr1, arr[i])
	}
	arr1 = append(arr1, arr[0:len(arr)-d]...)
	return arr1
}
