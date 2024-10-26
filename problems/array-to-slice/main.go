package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[:]
	slice = append(slice, 6, 7, 8, 9, 10)
	loop(slice)
}
func loop(slice []int) {
	for _, v := range slice {
		fmt.Println(v)
	}
}
