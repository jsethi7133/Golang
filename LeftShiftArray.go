package main

import "fmt"

func main() {
	arr := []int{64, 25, 12, 22, 11}
  d := 2
	fmt.Println(leftshift(arr, d))
}

func leftshift(arr []int, d int) []int {
	var arr1 []int
	for i := d; i < len(arr); i++ {
		arr1 = append(arr1, arr[i])
	}
	arr1 = append(arr1, arr[0:d]...)
	return arr1
}
